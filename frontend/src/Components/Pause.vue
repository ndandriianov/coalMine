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
    <span
        class="dot"
        :class="{ paused: isGameOnPause }"
    />
      <span>
      {{ isGameOnPause ? "Игра на паузе" : "Игра идёт" }}
    </span>
    </div>

    <button
        class="pause-button"
        @click="isGameOnPause ? resume() : pause()"
    >
      {{ isGameOnPause ? "Продолжить" : "Поставить на паузу" }}
    </button>
  </div>

</template>

<style scoped>
.pause-panel {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 12px 16px;
  border: 1px solid #ccc;
  border-radius: 8px;
  width: fit-content;
}

.status {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 500;
}

.dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: green;
}

.dot.paused {
  background: orange;
}

.pause-button {
  padding: 6px 12px;
  border-radius: 6px;
  border: none;
  cursor: pointer;
}

</style>