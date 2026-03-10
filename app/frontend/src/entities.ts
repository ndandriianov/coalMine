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

export const EquipmentType = {
  Pickaxe: 0,
  Ventilation: 1,
  Minecarts: 2,
} as const

export type EquipmentType =
  typeof EquipmentType[keyof typeof EquipmentType]


export const EquipmentLabels: Record<EquipmentType, string> = {
  [EquipmentType.Pickaxe]: "Кирка",
  [EquipmentType.Ventilation]: "Вентиляция",
  [EquipmentType.Minecarts]: "Вагонетки",
}