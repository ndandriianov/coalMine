<script setup lang="ts">
import {ref, onMounted, onUnmounted} from "vue";
import axios from "axios";

interface StartedResponse {
  Started: boolean;
}

const started = ref<boolean>(false);
const error = ref<string | null>(null)
const isLoading = ref<boolean>(false)

let intervalId: number | undefined;

async function fetchStartStatus() {
  try {
    isLoading.value = true;
    error.value = null;

    const response = await axios.get<StartedResponse>(
        "http://localhost:9091/mine/start"
    );

    started.value = response.data.Started;
  } catch (e) {
    error.value = "Ошибка при получении статуса";
  } finally {
    isLoading.value = false;
  }
}

async function startGame() {
  try {
    error.value = null;
    await axios.post("http://localhost:9091/mine/start");
  } catch (e) {
    error.value = "Ошибка при попытке начать игру";
  }
}

// lifecycle
onMounted(() => {
  fetchStartStatus();

  intervalId = window.setInterval(() => {
    fetchStartStatus();
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
    <div v-if="started">
      Игра начата
    </div>
    <div v-else>
      Игра не начата
    </div>
    <button v-if="!started" @click.prevent="startGame">
      Начать
    </button>
  </div>
</template>

<style scoped>

</style>