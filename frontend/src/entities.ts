export interface MinerInfo {
  Type: string
  EnergyLeft: number
  CoalPerExtraction: number
  CoalExtracted: number
  Started: boolean
  SleepTimeSeconds: number
}

export interface EquipmentInfo {
  Pickaxe: boolean
  Ventilation: boolean
  Minecarts: boolean
}