<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue"
import axios from "axios"

import type {EquipmentInfo} from "../entities.ts";


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
      <div>
        {{equipment}}
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>
