<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue";
import axios from "axios";

interface BalanceResponse {
  Balance: number;
}

const balance = ref<number>(0)
const error = ref<string | null>(null)
const isLoading = ref<boolean>(false)

let intervalId: number | undefined;

async function fetchBalance() {
  try {
    error.value = null;

    const response = await axios.get<BalanceResponse>(
        "http://localhost:9091/mine/balance"
    );

    balance.value = response.data.Balance;
  } catch (e) {
    error.value = "Ошибка при получении счётчика";
  } finally {
    isLoading.value = false;
  }
}

// lifecycle
onMounted(() => {
  fetchBalance();

  intervalId = window.setInterval(() => {
    fetchBalance();
  }, 500);
});

onUnmounted(() => {
  if (intervalId) {
    clearInterval(intervalId);
  }
});
</script>

<template>
  <div class="balance-widget" :class="{ 'is-loading': isLoading }">
    <div class="resource-icon">
      <div class="coal-crystal"></div>
    </div>

    <div class="balance-content">
      <span class="label">УГОЛЬ</span>
      <div class="value-wrapper">
        <span class="value">{{ balance.toLocaleString() }}</span>
        <span class="unit">ед.</span>
      </div>
    </div>

    <transition name="fade">
      <div v-if="error" class="balance-error" :title="error">!</div>
    </transition>
  </div>
</template>

<style scoped>
.balance-widget {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 10px 20px;
  background: #1a1a1a;
  border: 1px solid #333;
  border-radius: 12px;
  min-width: 160px;
  position: relative;
  transition: all 0.3s ease;
  user-select: none;
}

.balance-widget:hover {
  border-color: #444;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.4);
}

.resource-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  background: #262626;
  border-radius: 8px;
}

.coal-crystal {
  width: 18px;
  height: 18px;
  background: linear-gradient(135deg, #424242 0%, #000000 100%);
  transform: rotate(45deg);
  border-radius: 2px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
  border: 1px solid #555;
}

.balance-content {
  display: flex;
  flex-direction: column;
}

.label {
  font-size: 10px;
  font-weight: 800;
  color: #666;
  letter-spacing: 1.5px;
  line-height: 1;
  margin-bottom: 4px;
}

.value-wrapper {
  display: flex;
  align-items: baseline;
  gap: 4px;
}

.value {
  font-family: 'JetBrains Mono', 'Fira Code', monospace;
  font-size: 20px;
  font-weight: 700;
  color: #fff;
  line-height: 1;
}

.unit {
  font-size: 12px;
  color: #444;
  font-weight: 600;
}

.balance-error {
  position: absolute;
  top: -5px;
  right: -5px;
  width: 18px;
  height: 18px;
  background: #c62828;
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 11px;
  font-weight: bold;
  cursor: help;
}

.is-loading {
  opacity: 0.7;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>