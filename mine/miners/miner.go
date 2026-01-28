package miners

import "sync"

type Miner interface {
	Run(group *sync.WaitGroup)
	Info() MinerInfo
}
