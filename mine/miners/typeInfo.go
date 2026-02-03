package miners

type TypeInfo struct {
	Cost              int
	Energy            int
	CoalPerExtraction int
	SleepTimeSeconds  int
	Progress          int
}

func NewSmallMinerTypeInfo() TypeInfo {
	return TypeInfo{
		Cost:              5,
		Energy:            30,
		CoalPerExtraction: 1,
		SleepTimeSeconds:  3,
		Progress:          0,
	}
}

func NewMiddleMinerTypeInfo() TypeInfo {
	return TypeInfo{
		Cost:              50,
		Energy:            45,
		CoalPerExtraction: 3,
		SleepTimeSeconds:  2,
		Progress:          0,
	}
}

func NewStrongMinerTypeInfo() TypeInfo {
	return TypeInfo{
		Cost:              450,
		Energy:            60,
		CoalPerExtraction: 10,
		SleepTimeSeconds:  1,
		Progress:          3,
	}
}
