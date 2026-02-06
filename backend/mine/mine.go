package mine

import (
	"coalMine/helpers"
	"coalMine/mine/equipment"
	"coalMine/mine/errors"
	"coalMine/mine/miners"
	"coalMine/mine/pauseController"
	"coalMine/mine/resources"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	pickaxePrice     = 3000
	ventilationPrice = 15000
	minecartsPrice   = 50000
)

type Service struct {
	miners     []miners.Miner
	minersMtx  sync.RWMutex
	Balance    resources.Coal
	BalanceMtx sync.Mutex
	coalChan   chan resources.Coal

	producers *sync.WaitGroup
	consumers *sync.WaitGroup

	ctx       context.Context
	cancel    context.CancelFunc
	startOnce sync.Once
	started   atomic.Bool

	pickaxe      bool
	ventilation  bool
	minecarts    bool
	equipmentMtx sync.Mutex

	pc *pauseController.PauseController
}

func NewService() *Service {
	mineContext, mineCancel := context.WithCancel(context.Background())

	s := &Service{
		miners:   make([]miners.Miner, 0, 10),
		Balance:  0,
		coalChan: make(chan resources.Coal),

		producers: &sync.WaitGroup{},
		consumers: &sync.WaitGroup{},

		ctx:    mineContext,
		cancel: mineCancel,

		pc: pauseController.NewPauseController(),
	}

	s.started.Store(false)
	return s
}

func (s *Service) Start() {
	s.startOnce.Do(func() {
		s.started.Store(true)
		s.ctx, s.cancel = context.WithCancel(context.Background())
		s.producers = &sync.WaitGroup{}
		s.consumers = &sync.WaitGroup{}

		s.producers.Add(1)
		go PassiveIncome(s.ctx, s.producers, s.pc, s.coalChan)

		s.consumers.Add(1)
		go s.collectCoal()

		go func() {
			s.producers.Wait()
			close(s.coalChan)
		}()
	})
}

func (s *Service) Started() bool {
	return s.started.Load()
}

func (s *Service) Stop() {
	s.cancel()
	fmt.Println("Stopping the mine")

	s.producers.Wait()
	s.consumers.Wait()

	fmt.Println("Service stopped")
}

func (s *Service) Pause() {
	s.pc.Pause()
}

func (s *Service) Resume() {
	s.pc.Resume()
}

func (s *Service) IsOnPause() bool {
	return s.pc.IsOnPause()
}

// HireMiner adds a miner to the mine.
//
// Returns the miners id.
func (s *Service) HireMiner(minerType string) (int, error) {
	var miner miners.Miner
	switch minerType {
	case "small":
		miner = miners.NewSmallMiner(s.pc, s.coalChan)
	case "middle":
		miner = miners.NewMiddleMiner(s.pc, s.coalChan)
	case "strong":
		miner = miners.NewStrongMiner(s.pc, s.coalChan)
	default:
		return 0, errors.ErrInvalidMinerType
	}

	s.minersMtx.Lock()
	defer s.minersMtx.Unlock()
	s.miners = append(s.miners, miner)

	return len(s.miners), nil
}

func (s *Service) RunMiner(id int) bool {
	s.minersMtx.RLock()
	currentMiner, ok := helpers.SafeGet(id, s.miners)
	s.minersMtx.RUnlock()

	if !ok {
		return false
	}

	s.producers.Add(1)
	go currentMiner.Run(s.ctx, s.producers)

	return true
}

func (s *Service) RunAllNotStartedMiners() {
	s.minersMtx.RLock()
	for _, miner := range s.miners {
		if !miner.HasStarted() {
			s.producers.Add(1)
			go miner.Run(s.ctx, s.producers)
		}
	}
	s.minersMtx.RUnlock()
}

func (s *Service) collectCoal() {
	defer s.consumers.Done()

	for coal := range s.coalChan {
		s.BalanceMtx.Lock()
		s.Balance += coal
		s.BalanceMtx.Unlock()
	}

	fmt.Println("finished collecting coal")
}

// GetMiners returns info about all working miners.
//
// If type is empty string, GetMiners will return info for all types of miners.
func (s *Service) GetMiners(minerType string) map[int]miners.MinerInfo {
	collection := make(map[int]miners.MinerInfo)

	s.minersMtx.RLock()
	for id, miner := range s.miners {
		info := miner.Info()
		if info.Type == minerType || minerType == "" {
			collection[id] = info
		}
	}
	s.minersMtx.RUnlock()

	return collection
}

func (s *Service) BuyEquipment(equipmentType equipment.Type) error {
	s.BalanceMtx.Lock()
	defer s.BalanceMtx.Unlock()

	s.equipmentMtx.Lock()
	defer s.equipmentMtx.Unlock()

	switch equipmentType {
	case equipment.Pickaxe:
		if s.Balance >= pickaxePrice && !s.pickaxe {
			s.pickaxe = true
		} else if s.pickaxe {
			return errors.ErrEquipmentIsAlreadyBought
		} else {
			return errors.ErrNotEnoughCoal
		}

	case equipment.Ventilation:
		if s.Balance >= ventilationPrice && !s.ventilation {
			s.ventilation = true
		} else if s.ventilation {
			return errors.ErrEquipmentIsAlreadyBought
		} else {
			return errors.ErrNotEnoughCoal
		}

	case equipment.Minecarts:
		if s.Balance >= minecartsPrice && !s.minecarts {
			s.minecarts = true
		} else if s.minecarts {
			return errors.ErrEquipmentIsAlreadyBought
		} else {
			return errors.ErrNotEnoughCoal
		}
	}

	return nil
}

func (s *Service) GetEquipment() equipment.EquipmentInfo {
	s.equipmentMtx.Lock()
	defer s.equipmentMtx.Unlock()

	info := equipment.EquipmentInfo{
		Pickaxe:     s.pickaxe,
		Ventilation: s.ventilation,
		Minecarts:   s.minecarts,
	}

	return info
}
