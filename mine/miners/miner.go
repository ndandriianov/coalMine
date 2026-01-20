package miners

import (
	"coalMine/mine/resources"
	"context"
)

type Miner interface {
	Run(ctx context.Context) <-chan resources.Coal
	Info() MinerInfo
}
