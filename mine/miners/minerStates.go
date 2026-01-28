package miners

type MinerState int

const (
	NotStarted MinerState = iota
	Working
	Finished
	Pause
)
