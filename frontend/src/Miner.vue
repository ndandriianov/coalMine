<script setup lang="ts">
import MyButton from "./Components/UI/MyButton.vue";
import type {MinerInfo} from "./entities.ts";

defineProps<{
  miner: MinerInfo
  minerId: number
}>()
</script>

<template>
  <div>
    <strong>ID:</strong> {{ minerId }}<br/>
    <strong>Тип:</strong> {{ miner.Type }}<br/>
    <strong>Энергия:</strong> {{ miner.EnergyLeft }}<br/>
    <strong>Уголь за добычу:</strong> {{ miner.CoalPerExtraction }}<br/>
    <strong>Добыто угля:</strong> {{ miner.CoalExtracted }}<br/>
    <strong>Запущен:</strong> {{ miner.Started ? "Да" : "Нет" }}<br/>
    <strong>Время сна:</strong> {{ miner.SleepTimeSeconds }} сек

    <MyButton v-if="!miner.Started">
      Запустить
    </MyButton>
  </div>
</template>

<style scoped>

</style>