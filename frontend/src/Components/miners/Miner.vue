<script setup lang="ts">
import MyButton from "../UI/MyButton.vue";
import type {MinerInfo} from "../../entities.ts";
import axios from "axios";

const props = defineProps<{
  miner: MinerInfo
  minerId: number
}>()

async function StartMiner() {
  try {
    await axios.post("http://localhost:9091/mine/miner/start", null, {
      params: {
        id: props.minerId,
      }
    })
  } catch (e) {
    console.error("не удалось запустить шахтера:", e)
  }
}
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

    <MyButton v-if="!miner.Started" @click.prevent="StartMiner">
      Запустить
    </MyButton>
  </div>
</template>

<style scoped>

</style>