<script setup lang="ts">
import {ref, onMounted, onUnmounted, computed} from "vue"
import axios from "axios"

import type {MinerInfo} from "../entities.ts";
import Miner from "../Miner.vue";
import MyButton from "./UI/MyButton.vue";
import CollapseToggle from "./UI/CollapseToggle.vue";


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
    await fetchMiners()
  } catch (e) {
    console.error("не удалось запустить всех майнеров", e)
  }
}


const notStartedMiners = computed(() =>
    Object.entries(miners.value).filter(
        ([_, miner]) => !miner.Started
    )
)

const workingMiners = computed(() =>
    Object.entries(miners.value).filter(
        ([_, miner]) => miner.Started && miner.EnergyLeft > 0
    )
)

const exhaustedMiners = computed(() =>
    Object.entries(miners.value).filter(
        ([_, miner]) => miner.EnergyLeft === 0
    )
)

const isMinersOpen = ref(true)
const isNotStartedOpen = ref(true)
const isWorkingOpen = ref(true)
const isExhaustedOpen = ref(true)


function countByType(
    miners: Array<[string, MinerInfo]>
) {
  return miners.reduce(
      (acc, [_, miner]) => {
        acc[miner.Type] = (acc[miner.Type] ?? 0) + 1
        return acc
      },
      {} as Record<string, number>
  )
}

const notStartedCounts = computed(() =>
    countByType(notStartedMiners.value)
)

const workingCounts = computed(() =>
    countByType(workingMiners.value)
)

const exhaustedCounts = computed(() =>
    countByType(exhaustedMiners.value)
)

function formatCounts(counts: Record<string, number>) {
  const parts: string[] = []

  if (counts.small) parts.push(`small: ${counts.small}`)
  if (counts.middle) parts.push(`middle: ${counts.middle}`)
  if (counts.strong) parts.push(`strong: ${counts.strong}`)

  return parts.join(", ")
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

      <div v-if="notStartedMiners.length">
        <h3 class="section-title">
          <CollapseToggle
              :is-open="isNotStartedOpen"
              @click="isNotStartedOpen = !isNotStartedOpen"
          />
          Не запущены ({{ notStartedMiners.length }})

          <span v-if="!isNotStartedOpen" class="section-counts">
            — {{ formatCounts(notStartedCounts) }}
          </span>
        </h3>


        <ul v-show="isNotStartedOpen">
          <li v-for="[id, miner] in notStartedMiners" :key="id">
            <Miner :miner="miner" :miner-id="Number(id)"/>
          </li>
        </ul>
      </div>

      <div v-if="workingMiners.length">
        <h3 class="section-title">
          <CollapseToggle
              :is-open="isWorkingOpen"
              @click="isWorkingOpen = !isWorkingOpen"
          />
          Работают ({{ workingMiners.length }})

          <span v-if="!isWorkingOpen" class="section-counts">
            — {{ formatCounts(workingCounts) }}
          </span>
        </h3>


        <ul v-show="isWorkingOpen">
          <li v-for="[id, miner] in workingMiners" :key="id">
            <Miner :miner="miner" :miner-id="Number(id)"/>
          </li>
        </ul>
      </div>

      <div v-if="exhaustedMiners.length">
        <h3 class="section-title">
          <CollapseToggle
              :is-open="isExhaustedOpen"
              @click="isExhaustedOpen = !isExhaustedOpen"
          />
          Без энергии ({{ exhaustedMiners.length }})

          <span v-if="!isExhaustedOpen" class="section-counts">
            — {{ formatCounts(exhaustedCounts) }}
          </span>
        </h3>


        <ul v-show="isExhaustedOpen">
          <li v-for="[id, miner] in exhaustedMiners" :key="id">
            <Miner :miner="miner" :miner-id="Number(id)"/>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>


<style scoped>
.section-title {
  display: flex;
  align-items: center;
  user-select: none;
  margin: 16px 0 8px;
}

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

.miners-title {
  display: flex;
  align-items: center;
  user-select: none;
  margin-bottom: 12px;
}

.section-counts {
  font-size: 0.9em;
  color: #666;
  margin-left: 6px;
}

</style>
