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
  <div>
    <h2>Оборудование</h2>

    <div v-if="error">
      {{ error }}
    </div>

    <div v-else>
      <ul class="equipment-list">
        <li v-for="(value, name) in equipment" :key="name">
          <span class="label">{{ name }}</span>

          <div style="display: flex; flex-direction: column; align-items: center; gap: 5px">
            <span
                class="status"
                :class="{ on: value, off: !value }"
            >
              {{ value ? "Есть" : "Нет" }}
            </span>
            <MyButton v-if="!value" @click.prevent="() => buyEquipment(name)">
              Купить
            </MyButton>
          </div>
        </li>
      </ul>
    </div>

  </div>

  <div v-if="showModal" class="modal-overlay">
    <div class="modal">
      <h3>Поздравляем!</h3>
      <p>Всё оборудование куплено.</p>

      <MyButton @click="restartTheGame">
        Перезапустить игру
      </MyButton>
    </div>
  </div>

</template>

<style scoped>
.equipment-list {
  padding: 0;
  margin: 0;
}

.equipment-list li {
  list-style: none;
  display: flex;
  justify-content: space-between;
  padding: 8px 12px;
  border: 1px solid #ccc;
  border-radius: 6px;
  margin-bottom: 8px;
}

.label {
  font-weight: 500;
}

.status.on {
  color: green;
}

.status.off {
  color: red;
}

.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal {
  background: white;
  padding: 20px 24px;
  border-radius: 10px;
  min-width: 260px;
  text-align: center;
}

</style>
