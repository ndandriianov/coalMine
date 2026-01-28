package miners

import "sync"

type Miner interface {
	Run(group *sync.WaitGroup)
	Pause()
	Resume()
	Info() MinerInfo
}
