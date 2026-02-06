<script setup lang="ts">
import {ref, onMounted, onUnmounted, computed} from "vue"
import axios from "axios"

import type {MinerInfo} from "../entities.ts";
import Miner from "../Miner.vue";
import MyButton from "./UI/MyButton.vue";


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


const hasNotStarted = computed(() => {
  return Object.values(miners.value).some(miner => !miner.Started)
})

async function runAllMiners() {
  try {
    await axios.post("http://localhost:9091/mine/miner/start")
    fetchMiners() // обновляем состояние после запуска
  } catch (e) {
    console.error("не удалось запустить всех майнеров", e)
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

    <MyButton
        v-if="hasNotStarted"
        @click.prevent="runAllMiners"
        style="margin-bottom: 12px"
    >
      Запустить все
    </MyButton>

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
        <Miner :miner="miner" :miner-id="Number(id)" :key="id"/>
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
