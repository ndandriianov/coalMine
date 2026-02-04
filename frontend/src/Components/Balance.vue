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
  <div>Уголь {{ balance }}</div>
</template>

<style scoped>

</style>