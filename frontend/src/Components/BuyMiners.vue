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
    <h2 class="hire-title">Нанять шахтёра</h2>

    <div class="buttons-grid">
      <MyButton
          @click="hireMiner('small')"
          :disabled="loading"
          class="hire-btn small-type"
      >
        Маленький
      </MyButton>

      <MyButton
          @click="hireMiner('middle')"
          :disabled="loading"
          class="hire-btn middle-type"
      >
        Средний
      </MyButton>

      <MyButton
          @click="hireMiner('strong')"
          :disabled="loading"
          class="hire-btn strong-type"
      >
        Сильный
      </MyButton>
    </div>

    <div class="status-messages">
      <transition name="status-fade">
        <div v-if="loading" class="loading-state">
          <span class="spinner"></span>
          Отправка запроса...
        </div>
      </transition>

      <transition name="status-fade">
        <div v-if="error" class="error-msg">{{ error }}</div>
      </transition>

      <transition name="status-fade">
        <div v-if="result" class="success-msg">
          Шахтёр успешно нанят! ID: {{ result.id }}
        </div>
      </transition>
    </div>
  </div>
</template>

<style scoped>
.miner-hire {
  padding: 24px;
  background: #1a1a1a;
  border-radius: 16px;
  border: 1px solid #333;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  margin-bottom: 24px;
  max-width: 600px;
}

.hire-title {
  margin: 0 0 20px 0;
  font-size: 1.5rem;
  font-weight: 700;
  background: linear-gradient(to right, #ffffff, #999999);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.buttons-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(140px, 1fr));
  gap: 12px;
}

.hire-btn {
  width: 100%;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  font-size: 12px !important;
}

.small-type { background: #455a64 !important; }
.middle-type { background: #5e35b1 !important; }
.strong-type { background: #c62828 !important; }

.status-messages {
  margin-top: 20px;
  min-height: 24px;
}

.loading-state {
  color: #888;
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 14px;
}

.spinner {
  width: 16px;
  height: 16px;
  border: 2px solid rgba(255,255,255,0.1);
  border-top-color: #3949ab;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

.error-msg {
  color: #ef9a9a;
  background: rgba(239, 154, 154, 0.1);
  padding: 10px 14px;
  border-radius: 8px;
  border-left: 4px solid #ef9a9a;
  font-size: 14px;
}

.success-msg {
  color: #a5d6a7;
  background: rgba(165, 214, 167, 0.1);
  padding: 10px 14px;
  border-radius: 8px;
  border-left: 4px solid #a5d6a7;
  font-size: 14px;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.status-fade-enter-active,
.status-fade-leave-active {
  transition: all 0.3s ease;
}

.status-fade-enter-from,
.status-fade-leave-to {
  opacity: 0;
  transform: translateY(-5px);
}
</style>