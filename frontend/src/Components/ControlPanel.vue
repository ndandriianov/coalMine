<script setup lang="ts">

import StartStatus from "./StartStatus.vue";
import Pause from "./Pause.vue";
</script>

<template>
  <div class="game-controls-wrapper">
    <div class="controls-container">
      <StartStatus />
      <div class="divider"></div>
      <Pause/>
    </div>
  </div>
</template>

<style scoped>
.game-controls-wrapper {
  padding: 16px;
  background: #1a1a1a;
  border-radius: 16px;
  border: 1px solid #333;
  display: inline-block;
  margin-bottom: 24px;
}

.controls-container {
  display: flex;
  align-items: center;
  gap: 24px;
}

.divider {
  width: 1px;
  height: 32px;
  background: #333;
}
</style>