<template>
  <button class="btn">
    <slot></slot>
  </button>
</template>

<style scoped>
.btn {
  padding: 10px 20px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;

  background: #3949ab;
  color: #ffffff;

  border: none;
  border-radius: 8px;

  box-shadow: 0 4px 12px rgba(57, 73, 171, 0.3);

  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.btn:hover {
  background: #5c6bc0;
  transform: translateY(-1px);
  box-shadow: 0 6px 16px rgba(57, 73, 171, 0.45);
}

.btn:active {
  background: #303f9f;
  transform: translateY(1px);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.btn:disabled {
  background: #333;
  color: #666;
  cursor: not-allowed;
  box-shadow: none;
  transform: none;
}
</style>