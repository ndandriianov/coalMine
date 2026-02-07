<script setup lang="ts">
import {ref, computed} from "vue"
import type { MinerInfo } from "../../entities.ts"
import Miner from "./Miner.vue"
import CollapseToggle from "../UI/CollapseToggle.vue"

const props = defineProps<{
  title: string
  miners: Array<[string, MinerInfo]>
}>()

const isOpen = ref(true)

function countByType(miners: Array<[string, MinerInfo]>) {
  return miners.reduce((acc, [_, miner]) => {
    acc[miner.Type] = (acc[miner.Type] ?? 0) + 1
    return acc
  }, {} as Record<string, number>)
}

const counts = computed(() => countByType(props.miners))

function formatCounts(counts: Record<string, number>) {
  const parts: string[] = []
  if (counts.small) parts.push(`small: ${counts.small}`)
  if (counts.middle) parts.push(`middle: ${counts.middle}`)
  if (counts.strong) parts.push(`strong: ${counts.strong}`)
  return parts.join(", ")
}
</script>

<template>
  <div v-if="miners.length">
    <h3 class="section-title">
      <CollapseToggle
          :is-open="isOpen"
          @click="isOpen = !isOpen"
      />
      {{ title }} ({{ miners.length }})

      <span v-if="!isOpen" class="section-counts">
        — {{ formatCounts(counts) }}
      </span>
    </h3>

    <ul v-show="isOpen">
      <li v-for="[id, miner] in miners" :key="id">
        <Miner :miner="miner" :miner-id="Number(id)" />
      </li>
    </ul>
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

.section-counts {
  font-size: 0.9em;
  color: #666;
  margin-left: 6px;
}
</style>
