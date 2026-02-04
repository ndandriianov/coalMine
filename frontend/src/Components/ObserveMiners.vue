<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue"
import axios from "axios"


interface MinerInfo {
  Type: string
  EnergyLeft: number
  CoalPerExtraction: number
  CoalExtracted: number
  Started: boolean
  SleepTimeSeconds: number
}

type MinersResponse = Record<number, MinerInfo>


const miners = ref<MinersResponse>({})
const isLoading = ref(false)
const error = ref<string | null>(null)

let intervalId: number | undefined;


async function fetchMiners() {
  try {
    isLoading.value = true
    error.value = null

    const response = await axios.get<MinersResponse>(
        "http://localhost:9091/mine/miner"
    )

    miners.value = response.data
  } catch (e) {
    error.value = "Ошибка при загрузке майнеров"
  } finally {
    isLoading.value = false
  }
}


// lifecycle
onMounted(() => {
  fetchMiners();

  intervalId = window.setInterval(() => {
    fetchMiners();
  }, 500);
});

onUnmounted(() => {
  if (intervalId) {
    clearInterval(intervalId);
  }
});
</script>

<template>
  <div>
    <h2>Майнеры</h2>

    <div v-if="isLoading && Object.keys(miners).length === 0">
      Загрузка...
    </div>

    <div v-else-if="error">
      {{ error }}
    </div>

    <div v-else-if="Object.keys(miners).length === 0">
      Майнеров нет
    </div>

    <ul v-else>
      <li v-for="(miner, id) in miners" :key="id">
        <strong>ID:</strong> {{ id }}<br />
        <strong>Тип:</strong> {{ miner.Type }}<br />
        <strong>Энергия:</strong> {{ miner.EnergyLeft }}<br />
        <strong>Уголь за добычу:</strong> {{ miner.CoalPerExtraction }}<br />
        <strong>Добыто угля:</strong> {{ miner.CoalExtracted }}<br />
        <strong>Запущен:</strong> {{ miner.Started ? "Да" : "Нет" }}<br />
        <strong>Время сна:</strong> {{ miner.SleepTimeSeconds }} сек
      </li>
    </ul>
  </div>
</template>

<style scoped>
ul {
  padding: 0;
}

li {
  list-style: none;
  padding: 12px;
  margin-bottom: 12px;
  border: 1px solid #ccc;
  border-radius: 6px;
}
</style>
