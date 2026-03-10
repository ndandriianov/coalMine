<script setup lang="ts">
defineProps<{
  isOpen: boolean
}>()
</script>

<template>
  <button class="toggle">
    <span class="icon">
      {{ isOpen ? "▼" : "▶" }}
    </span>
  </button>
</template>

<style scoped>
.toggle {
  background: none;
  border: none;
  cursor: pointer;
  padding: 0;
  margin-right: 6px;
  font-size: 14px;
}

.icon {
  display: inline-block;
  width: 14px;
  text-align: center;
  color: #1a4185;
}
</style>
