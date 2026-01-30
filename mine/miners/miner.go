package miners

import (
	"context"
	"sync"
)

type Miner interface {
	Run(ctx context.Context, group *sync.WaitGroup)
	Info() MinerInfo
	HasStarted() bool
}
