package miners

import (
	"coalMine/mine/resources"
)

type MinerInfo struct {
	Type              string
	EnergyLeft        int
	CoalPerExtraction resources.Coal
	CoalExtracted     int
	Started           bool
	SleepTimeSeconds  int
}
