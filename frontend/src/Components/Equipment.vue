<script setup lang="ts">
import {ref, onMounted, onUnmounted, computed, watch} from "vue"
import axios from "axios"

import type {EquipmentInfo} from "../entities.ts";
import MyButton from "./UI/MyButton.vue";


const equipment = ref<EquipmentInfo>()
const isLoading = ref(false)
const error = ref<string | null>(null)

let intervalId: number | undefined;


const allBought = computed(() => {
  if (!equipment.value) return false
  return Object.values(equipment.value).every((value) => value === true)
})
const showModal = ref(false)
watch(allBought, (newValue) => {
  if (newValue) {
    showModal.value = true
  }
})


async function fetchEquipment() {
  try {
    isLoading.value = true
    error.value = null

    const response = await axios.get<EquipmentInfo>(
        "http://localhost:9091/mine/equipment"
    )

    equipment.value = response.data
  } catch (e) {
    error.value = "Ошибка при загрузке оборудования"
  } finally {
    isLoading.value = false
  }
}

async function buyEquipment(equipment: string) {
  try {
    await axios.post("http://localhost:9091/mine/equipment", {
      Equipment: equipment
    })
  } catch (e) {
    console.error("не удалось купить оборудование", e)
  }
}

async function restartTheGame() {
  try {
    await axios.post("http://localhost:9091/mine/restart")
    showModal.value = false
  } catch (e) {
    console.error("не удалось перезапустить игру", e)
  }
}


// lifecycle
onMounted(() => {
  fetchEquipment();

  intervalId = window.setInterval(() => {
    fetchEquipment();
  }, 500);
});

onUnmounted(() => {
  if (intervalId) {
    clearInterval(intervalId);
  }
});
</script>

<template>
  <div class="equipment-container">
    <h2 class="equipment-title">Оборудование</h2>

    <div v-if="error" class="error-banner">
      {{ error }}
    </div>

    <div v-else class="equipment-grid">
      <div
          v-for="(isBought, name) in equipment"
          :key="name"
          class="equipment-card"
          :class="{ 'is-owned': isBought }"
      >
        <div class="card-header">
          <span class="equipment-name">{{ name }}</span>
          <div class="indicator" :class="{ 'active': isBought }"></div>
        </div>

        <div class="card-body">
          <span class="status-badge" :class="isBought ? 'owned' : 'missing'">
            {{ isBought ? "В наличии" : "Отсутствует" }}
          </span>

          <MyButton
              v-if="!isBought"
              @click.prevent="() => buyEquipment(name as string)"
              class="buy-button"
          >
            Приобрести
          </MyButton>
        </div>
      </div>
    </div>

    <transition name="modal-fade">
      <div v-if="showModal" class="modal-overlay">
        <div class="modal-content">
          <div class="success-icon">🏆</div>
          <h3 class="modal-title">Поздравляем!</h3>
          <p class="modal-text">Весь комплект оборудования собран. Шахта готова к работе на полную мощность.</p>

          <MyButton @click="restartTheGame" class="restart-btn">
            Перезапустить игру
          </MyButton>
        </div>
      </div>
    </transition>
  </div>
</template>

<style scoped>
.equipment-container {
  padding: 8px;
}

.equipment-title {
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 20px;
  color: #e0e0e0;
}

.equipment-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 16px;
}

/* Карточки оборудования */
.equipment-card {
  background: #1a1a1a;
  border: 1px solid #333;
  border-radius: 12px;
  padding: 16px;
  transition: all 0.3s ease;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.equipment-card.is-owned {
  border-color: rgba(57, 73, 171, 0.4);
  background: linear-gradient(145deg, #1a1a1a, #1e202e);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.equipment-name {
  font-weight: 600;
  color: #fff;
  text-transform: capitalize;
}

.indicator {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #444;
  box-shadow: inset 0 0 4px #000;
}

.indicator.active {
  background: #00e676;
  box-shadow: 0 0 10px rgba(0, 230, 118, 0.5);
}

.card-body {
  display: flex;
  flex-direction: column;
  gap: 10px;
  align-items: center;
}

.status-badge {
  font-size: 11px;
  font-weight: 700;
  padding: 2px 8px;
  border-radius: 4px;
  text-transform: uppercase;
}

.status-badge.owned {
  color: #00e676;
  background: rgba(0, 230, 118, 0.1);
}

.status-badge.missing {
  color: #ff5252;
  background: rgba(255, 82, 82, 0.1);
}

.buy-button {
  width: 100%;
  font-size: 12px !important;
  padding: 8px !important;
}

/* Модальное окно */
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.85);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background: #1e1e1e;
  border: 1px solid #3949ab;
  padding: 32px;
  border-radius: 20px;
  max-width: 340px;
  text-align: center;
  box-shadow: 0 0 40px rgba(57, 73, 171, 0.25);
}

.success-icon {
  font-size: 48px;
  margin-bottom: 16px;
}

.modal-title {
  color: #fff;
  margin-bottom: 8px;
}

.modal-text {
  color: #aaa;
  font-size: 14px;
  margin-bottom: 24px;
  line-height: 1.5;
}

.restart-btn {
  width: 100%;
  background: #c62828 !important;
}

/* Анимации */
.modal-fade-enter-active, .modal-fade-leave-active {
  transition: opacity 0.3s ease;
}
.modal-fade-enter-from, .modal-fade-leave-to {
  opacity: 0;
}

.error-banner {
  background: rgba(255, 82, 82, 0.1);
  color: #ff5252;
  padding: 12px;
  border-radius: 8px;
  border: 1px solid rgba(255, 82, 82, 0.3);
}
</style>