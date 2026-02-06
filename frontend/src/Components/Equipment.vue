<script setup lang="ts">
import {ref, onMounted, onUnmounted} from "vue"
import axios from "axios"

import type {EquipmentInfo} from "../entities.ts";
import MyButton from "./UI/MyButton.vue";


type EquipmentResponse = Record<number, EquipmentInfo>


const equipment = ref<EquipmentResponse>({})
const isLoading = ref(false)
const error = ref<string | null>(null)

let intervalId: number | undefined;


async function fetchEquipment() {
  try {
    isLoading.value = true
    error.value = null

    const response = await axios.get<EquipmentResponse>(
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

    <div v-if="isLoading && Object.keys(equipment).length === 0">
      Загрузка...
    </div>

    <div v-else-if="error">
      {{ error }}
    </div>

    <div v-else-if="Object.keys(equipment).length === 0">
      Оборудования нет
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

</style>
