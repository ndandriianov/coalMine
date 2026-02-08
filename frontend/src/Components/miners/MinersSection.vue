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
  <div v-if="miners.length" class="miners-section">
    <header class="section-header" @click="isOpen = !isOpen">
      <div class="title-group">
        <CollapseToggle :is-open="isOpen" />
        <h3 class="section-title">
          {{ title }}
          <span class="count-badge">{{ miners.length }}</span>
        </h3>
      </div>

      <transition name="fade">
        <span v-if="!isOpen" class="section-counts">
          {{ formatCounts(counts) }}
        </span>
      </transition>
    </header>

    <transition name="expand">
      <ul v-show="isOpen" class="miners-grid">
        <li v-for="[id, miner] in miners" :key="id" class="miner-item">
          <Miner :miner="miner" :miner-id="Number(id)" />
        </li>
      </ul>
    </transition>
  </div>
</template>

<style scoped>
.miners-section {
  margin-bottom: 32px;
  font-family: 'Inter', -apple-system, sans-serif;
}

.section-header {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  gap: 16px;
  cursor: pointer;
  padding: 8px 12px;
  border-radius: 8px;
  transition: background 0.2s;
  user-select: none;
}

.section-header:hover {
  background: rgba(255, 255, 255, 0.05);
}

.title-group {
  display: flex;
  align-items: center;
  gap: 12px;
}

.section-title {
  margin: 0;
  font-size: 1.2rem;
  font-weight: 600;
  color: #888;
  display: flex;
  align-items: center;
  gap: 8px;
}

.count-badge {
  font-size: 0.8rem;
  background: #333;
  color: #888;
  padding: 2px 8px;
  border-radius: 12px;
}

.section-counts {
  font-size: 0.85rem;
  color: #888;
  letter-spacing: 0.5px;
  border-left: 2px solid #444;
  padding-left: 12px;
}

.miners-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 20px;
  padding: 16px 0;
  margin: 0;
  list-style: none;
}

.miner-item {
  display: flex;
  justify-content: center;
}

/* Анимации */
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

.expand-enter-active, .expand-leave-active {
  transition: all 0.3s ease-out;
  max-height: 2000px;
  overflow: hidden;
}
.expand-enter-from, .expand-leave-to {
  max-height: 0;
  opacity: 0;
  transform: translateY(-10px);
}
</style>