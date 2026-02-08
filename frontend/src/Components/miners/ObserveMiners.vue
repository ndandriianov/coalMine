<script setup lang="ts">
import {ref, onMounted, onUnmounted, computed} from "vue"
import axios from "axios"

import type { MinerInfo } from "../../entities.ts"
import MyButton from "../UI/MyButton.vue"
import CollapseToggle from "../UI/CollapseToggle.vue"
import MinersSection from "./MinersSection.vue"

type MinersResponse = Record<number, MinerInfo>

const miners = ref<MinersResponse>({})
let intervalId: number | undefined

async function fetchMiners() {
  const response = await axios.get<MinersResponse>(
      "http://localhost:9091/mine/miner"
  )
  miners.value = response.data
}

const hasNotStarted = computed(() =>
    Object.values(miners.value).some(m => !m.Started)
)

async function runAllMiners() {
  await axios.post("http://localhost:9091/mine/miner/start")
  await fetchMiners()
}

const notStartedMiners = computed(() =>
    Object.entries(miners.value).filter(([_, m]) => !m.Started)
)

const workingMiners = computed(() =>
    Object.entries(miners.value).filter(
        ([_, m]) => m.Started && m.EnergyLeft > 0
    )
)

const exhaustedMiners = computed(() =>
    Object.entries(miners.value).filter(
        ([_, m]) => m.EnergyLeft === 0
    )
)

const isMinersOpen = ref(true)

onMounted(() => {
  fetchMiners()
  intervalId = window.setInterval(fetchMiners, 500)
})

onUnmounted(() => {
  if (intervalId) clearInterval(intervalId)
})
</script>

<template>
  <div class="miners-container">
    <header class="miners-main-header">
      <div class="title-wrapper" @click="isMinersOpen = !isMinersOpen">
        <CollapseToggle :is-open="isMinersOpen" />
        <h2 class="miners-title">Майнеры</h2>
      </div>

      <transition name="fade">
        <MyButton
            v-if="hasNotStarted && isMinersOpen"
            class="run-all-btn"
            @click.prevent="runAllMiners"
        >
          <span>▶</span> Запустить все
        </MyButton>
      </transition>
    </header>

    <transition
        name="collapse"
        @enter="el => (el as HTMLElement).style.height = (el as HTMLElement).scrollHeight + 'px'"
        @after-enter="el => (el as HTMLElement).style.height = 'auto'"
        @leave="el => (el as HTMLElement).style.height = (el as HTMLElement).scrollHeight + 'px'"
        @before-leave="el => (el as HTMLElement).offsetHeight"
    >
      <div v-if="isMinersOpen" class="content-clipper">
        <div class="sections-wrapper">
          <MinersSection title="Не запущены" :miners="notStartedMiners" />
          <MinersSection title="Работают" :miners="workingMiners" />
          <MinersSection title="Без энергии" :miners="exhaustedMiners" />
        </div>
      </div>
    </transition>
  </div>
</template>

<style scoped>
.miners-container {
  padding: 20px;
  background: #121212;
  color: #e0e0e0;
  border-radius: 10px;
}

.miners-main-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 1px solid #333;
}

.title-wrapper {
  display: flex;
  align-items: center;
  gap: 12px;
  cursor: pointer;
}

.miners-title {
  font-size: 1.8rem;
  font-weight: 700;
  margin: 0;
  background: linear-gradient(to right, #fff, #888);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.run-all-btn {
  background: #2e7d32 !important;
  box-shadow: 0 4px 14px rgba(46, 125, 50, 0.4) !important;
}

.content-clipper {
  overflow: hidden;
  transition: height 0.35s cubic-bezier(0.4, 0, 0.2, 1), opacity 0.3s ease;
}

.sections-wrapper {
  display: flex;
  flex-direction: column;
  gap: 8px;
  padding-bottom: 20px;
}

/* Анимация схлопывания */
.collapse-enter-from,
.collapse-leave-to {
  height: 0 !important;
  opacity: 0;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.2s ease;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>