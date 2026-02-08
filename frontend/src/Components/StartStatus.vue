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
  <div class="start-panel">
    <div class="game-status" :class="{ 'is-active': started }">
      <div class="status-indicator"></div>
      <span class="status-label">
        {{ started ? "СЕРВЕР АКТИВЕН" : "СЕРВЕР ВЫКЛЮЧЕН" }}
      </span>
    </div>

    <MyButton
        v-if="!started"
        @click.prevent="startGame"
        class="start-main-btn"
    >
      ЗАПУСТИТЬ ИГРУ
    </MyButton>
  </div>
</template>

<style scoped>
.start-panel {
  display: flex;
  align-items: center;
  gap: 16px;
}

.game-status {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 6px 12px;
  background: rgba(255, 0, 0, 0.1);
  border-radius: 6px;
  border: 1px solid rgba(255, 0, 0, 0.2);
}

.game-status.is-active {
  background: rgba(0, 255, 100, 0.05);
  border-color: rgba(0, 255, 100, 0.2);
}

.status-indicator {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: #f44336;
}

.is-active .status-indicator {
  background: #00ff6a;
  box-shadow: 0 0 8px #00ff6a;
}

.status-label {
  font-size: 11px;
  font-weight: 700;
  color: #f44336;
}

.is-active .status-label {
  color: #00ff6a;
}

.start-main-btn {
  background: #3949ab !important;
  height: 36px;
  padding: 0 20px !important;
}
</style>