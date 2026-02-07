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
  <div class="miner-card">
    <div class="header">
      <span class="miner-id">Шахтёр #{{ minerId }}</span>
      <span
          class="status"
          :class="{ started: miner.Started, stopped: !miner.Started }"
      >
        {{ miner.Started ? "Работает" : "Остановлен" }}
      </span>
    </div>

    <div class="info">
      <div><span>Тип</span><span>{{ miner.Type }}</span></div>
      <div><span>Энергия</span><span>{{ miner.EnergyLeft }}</span></div>
      <div><span>Уголь за добычу</span><span>{{ miner.CoalPerExtraction }}</span></div>
      <div><span>Добыто угля</span><span>{{ miner.CoalExtracted }}</span></div>
      <div><span>Время сна</span><span>{{ miner.SleepTimeSeconds }} сек</span></div>
    </div>

    <MyButton
        v-if="!miner.Started"
        class="start-btn"
        @click.prevent="StartMiner"
    >
      ▶ Запустить
    </MyButton>
  </div>
</template>


<style scoped>
.miner-card {
  width: 320px;
  padding: 16px;
  border-radius: 12px;
  background: #1e1e1e;
  color: #fff;
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.35);
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.miner-id {
  font-weight: 600;
  font-size: 16px;
}

.status {
  padding: 4px 10px;
  border-radius: 999px;
  font-size: 12px;
  font-weight: 600;
}

.status.started {
  background: #1b5e20;
  color: #a5d6a7;
}

.status.stopped {
  background: #5f2120;
  color: #ef9a9a;
}

.info {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 10px 12px;
  font-size: 14px;
}

.info div {
  display: flex;
  flex-direction: column;
}

.info span:first-child {
  opacity: 0.6;
  font-size: 12px;
}

.start-btn {
  margin-top: 8px;
  align-self: center;
}
</style>
