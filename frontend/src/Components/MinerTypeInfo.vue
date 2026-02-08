<script setup lang="ts">
import { ref, onMounted } from "vue";
import axios from "axios";

const props = defineProps<{
  type: 'small' | 'middle' | 'strong'
}>();

interface TypeInfo {
  Cost: number;
  Energy: number;
  CoalPerExtraction: number;
  SleepTimeSeconds: number;
  Progress: number;
}

const info = ref<TypeInfo | null>(null);
const loading = ref(true);
const error = ref(false);

const typeLabels = {
  small: 'Маленький',
  middle: 'Средний',
  strong: 'Сильный'
};

async function fetchTypeInfo() {
  try {
    const response = await axios.get<TypeInfo>(
        `http://localhost:9091/mine/miner/info`,
        { params: { type: props.type } }
    );
    info.value = response.data;
  } catch (e) {
    error.value = true;
  } finally {
    loading.value = false;
  }
}

onMounted(fetchTypeInfo);
</script>

<template>
  <div class="type-info-card" :class="type">
    <div v-if="loading" class="skeleton">Загрузка характеристик...</div>

    <div v-else-if="info" class="info-content">
      <div class="info-header">
        <span class="type-name">{{ typeLabels[type] }}</span>
        <span class="cost">{{ info.Cost }}</span>
      </div>

      <div class="stats-grid">
        <div class="stat-item">
          <span class="stat-label">Энергия</span>
          <span class="stat-value">{{ info.Energy }}</span>
        </div>
        <div class="stat-item">
          <span class="stat-label">Добыча</span>
          <span class="stat-value">+{{ info.CoalPerExtraction }}</span>
        </div>
        <div class="stat-item">
          <span class="stat-label">Сон</span>
          <span class="stat-value">{{ info.SleepTimeSeconds }}с</span>
        </div>
        <div class="stat-item">
          <span class="stat-label">Прогресс</span>
          <span class="stat-value">{{ info.Progress }}%</span>
        </div>
      </div>
    </div>

    <div v-else class="error">!</div>
  </div>
</template>

<style scoped>
.type-info-card {
  background: #222;
  border: 1px solid #333;
  border-radius: 12px;
  padding: 12px 16px;
  min-width: 200px;
  transition: transform 0.2s;
}

.type-info-card:hover {
  border-color: #444;
  transform: translateY(-2px);
}

.info-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  border-bottom: 1px solid #333;
  padding-bottom: 8px;
}

.type-name {
  font-weight: 700;
  font-size: 0.9rem;
  text-transform: uppercase;
  color: #aaa;
}

/* Цветовые акценты для типов */
.small .type-name { color: #81c784; }
.middle .type-name { color: #64b5f6; }
.strong .type-name { color: #e57373; }

.cost {
  font-weight: 600;
  color: #ffd54f;
  font-size: 0.9rem;
}

.stats-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 8px 12px;
}

.stat-item {
  display: flex;
  flex-direction: column;
}

.stat-label {
  font-size: 10px;
  color: #666;
  text-transform: uppercase;
}

.stat-value {
  font-size: 13px;
  font-weight: 600;
  color: #eee;
}

.skeleton {
  font-size: 12px;
  color: #444;
  text-align: center;
  padding: 10px;
}

.error {
  color: #ff5252;
  text-align: center;
}
</style>