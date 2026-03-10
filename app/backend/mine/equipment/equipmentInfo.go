package equipment

type EquipmentInfo struct {
	Pickaxe     EquipmentItem
	Ventilation EquipmentItem
	Minecarts   EquipmentItem
}

type EquipmentItem struct {
	IsBought bool
	Price    int
}
