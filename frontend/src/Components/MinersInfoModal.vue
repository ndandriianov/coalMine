<script setup lang="ts">
import { ref } from "vue";
import MyButton from "./UI/MyButton.vue";
import MinerTypeInfo from "./MinerTypeInfo.vue";

const isModalOpen = ref(false);

function closeModal() {
  isModalOpen.value = false;
}
</script>

<template>
  <div class="miners-info-wrapper">
    <MyButton @click="isModalOpen = true" class="info-trigger-btn">
      <span class="icon">ℹ</span> Характеристики типов
    </MyButton>

    <transition name="modal-fade">
      <div v-if="isModalOpen" class="modal-overlay" @click.self="closeModal">
        <div class="modal-content">
          <header class="modal-header">
            <h3>Справочник шахтёров</h3>
            <button class="close-btn" @click="closeModal">&times;</button>
          </header>

          <div class="info-grid">
            <MinerTypeInfo type="small" />
            <MinerTypeInfo type="middle" />
            <MinerTypeInfo type="strong" />
          </div>

          <div class="modal-footer">
            <MyButton @click="closeModal">Закрыть</MyButton>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>

<style scoped>
.miners-info-wrapper {
  display: inline-block;
}

.info-trigger-btn {
  background: #455a64 !important;
  height: 36px;
  font-size: 13px !important;
}

.icon {
  margin-right: 6px;
  font-weight: bold;
}

.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.8);
  backdrop-filter: blur(6px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 2000;
  padding: 20px;
}

.modal-content {
  background: #1a1a1a;
  border: 1px solid #333;
  border-radius: 20px;
  width: 100%;
  max-width: 800px;
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.5);
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.modal-header {
  padding: 20px 24px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #333;
}

.modal-header h3 {
  margin: 0;
  color: #fff;
  font-size: 1.4rem;
}

.close-btn {
  background: none;
  border: none;
  color: #666;
  font-size: 28px;
  cursor: pointer;
  line-height: 1;
  transition: color 0.2s;
}

.close-btn:hover {
  color: #fff;
}

.info-grid {
  padding: 24px;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

.modal-footer {
  padding: 16px 24px;
  border-top: 1px solid #333;
  display: flex;
  justify-content: flex-end;
  background: #1e1e1e;
}

.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.3s ease;
}

.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}

@media (max-width: 800px) {
  .info-grid {
    grid-template-columns: 1fr;
    max-height: 60vh;
    overflow-y: auto;
  }

  .modal-content {
    max-width: 100%;
  }
}
</style>