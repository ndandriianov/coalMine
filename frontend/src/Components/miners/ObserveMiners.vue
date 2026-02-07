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
  <div>
    <h2 class="miners-title">
      <CollapseToggle
          :is-open="isMinersOpen"
          @click="isMinersOpen = !isMinersOpen"
      />
      Майнеры
    </h2>

    <div v-show="isMinersOpen">
      <MyButton
          v-if="hasNotStarted"
          @click.prevent="runAllMiners"
          style="margin-bottom: 12px"
      >
        Запустить все
      </MyButton>

      <MinersSection
          title="Не запущены"
          :miners="notStartedMiners"
      />

      <MinersSection
          title="Работают"
          :miners="workingMiners"
      />

      <MinersSection
          title="Без энергии"
          :miners="exhaustedMiners"
      />
    </div>
  </div>
</template>

<style scoped>
.miners-title {
  display: flex;
  align-items: center;
  user-select: none;
  margin-bottom: 12px;
}
</style>
