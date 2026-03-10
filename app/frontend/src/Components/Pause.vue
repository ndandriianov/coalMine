<script setup lang="ts">
import axios from "axios";
import {ref, onMounted, onUnmounted} from "vue"

interface PauseDto {
  IsOnPause: boolean
}

const isGameOnPause = ref<boolean>(false);
let intervalId: number | undefined;

async function pause() {
  try {
    await axios.post("http://localhost:9091/mine/pause")
  } catch (e) {
    console.error("не удалось поставить игру на паузу", e)
  }
}

async function resume() {
  try {
    await axios.post("http://localhost:9091/mine/resume")
  } catch (e) {
    console.error("не удалось снять игру с паузы", e)
  }
}

async function fetchPause() {
  try {
    const response = await axios.get<PauseDto>(
        "http://localhost:9091/mine/pause"
    )
    isGameOnPause.value = response.data.IsOnPause
  } catch (e) {
    console.error(e)
  }
}


// lifecycle
onMounted(() => {
  fetchPause();

  intervalId = window.setInterval(() => {
    fetchPause();
  }, 500);
});

onUnmounted(() => {
  if (intervalId) {
    clearInterval(intervalId);
  }
});
</script>

<template>
  <div class="pause-panel">
    <div class="status">
      <span class="dot" :class="{ paused: isGameOnPause }" />
      <span class="status-text">
        {{ isGameOnPause ? "ПАУЗА" : "В ПРОЦЕССЕ" }}
      </span>
    </div>

    <button
        class="action-button"
        :class="{ 'resume-mode': isGameOnPause }"
        @click="isGameOnPause ? resume() : pause()"
    >
      {{ isGameOnPause ? "Продолжить" : "Пауза" }}
    </button>
  </div>
</template>

<style scoped>
.pause-panel {
  display: flex;
  align-items: center;
  gap: 16px;
}

.status {
  display: flex;
  align-items: center;
  gap: 10px;
}

.status-text {
  font-size: 12px;
  font-weight: 800;
  letter-spacing: 1px;
  color: #888;
}

.dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #4caf50;
  box-shadow: 0 0 10px #4caf50;
  transition: all 0.3s ease;
}

.dot.paused {
  background: #ff9800;
  box-shadow: 0 0 10px #ff9800;
}

.action-button {
  padding: 8px 16px;
  border-radius: 8px;
  border: 1px solid #444;
  background: transparent;
  color: #eee;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.action-button:hover {
  background: #333;
  border-color: #666;
}

.action-button.resume-mode {
  border-color: #ff9800;
  color: #ff9800;
}
</style>