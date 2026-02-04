<script setup lang="ts">
import { ref } from "vue";
import axios from "axios";
import MyButton from "./UI/MyButton.vue";

interface HireMinerResponse {
  id: number;
}

const loading = ref(false);
const error = ref<string | null>(null);
const result = ref<HireMinerResponse | null>(null);

async function hireMiner(minerType: string) {
  if (minerType === "") return;

  loading.value = true;
  error.value = null;
  result.value = null;

  try {
    const response = await axios.post<HireMinerResponse>(
        "http://localhost:9091/mine/miner/hire",
        null, // тело POST пустое, данные через query
        {
          params: { type: minerType },
          withCredentials: true, // если сервер использует cookies
        }
    );

    result.value = response.data;
  } catch (e: any) {
    if (e.response && e.response.data) {
      error.value = e.response.data;
    } else {
      error.value = "Ошибка при запросе к серверу";
    }
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <div class="miner-hire">
    <h2>Нанять шахтёра</h2>

    <div class="buttons">
      <MyButton
          @click="hireMiner('small')"
          :disabled="loading"
      >
        Маленький
      </MyButton>

      <MyButton
          @click="hireMiner('middle')"
          :disabled="loading"
      >
        Средний
      </MyButton>

      <MyButton
          @click="hireMiner('strong')"
          :disabled="loading"
      >
        Сильный
      </MyButton>
    </div>

    <div v-if="loading">Отправка запроса...</div>
    <div v-if="error" class="error">{{ error }}</div>
    <div v-if="result" class="success">
      Шахтёр успешно нанят! ID: {{ result.id }}
    </div>
  </div>
</template>


<style scoped>

</style>