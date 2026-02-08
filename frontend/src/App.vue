<script setup lang="ts">
import BuyMiners from "./Components/BuyMiners.vue";
import Balance from "./Components/Balance.vue";
import ObserveMiners from "./Components/miners/ObserveMiners.vue";
import Equipment from "./Components/Equipment.vue";
import ControlPanel from "./Components/ControlPanel.vue";

</script>

<template>
  <div class="page-background">
    <div class="app-container">
      <div class="top-row">
        <BuyMiners />
        <div class="controls-column">
          <ControlPanel />
          <Balance />
        </div>
      </div>

      <div class="content-body">
        <ObserveMiners />
        <Equipment />
      </div>
    </div>
  </div>
</template>

<style>
body {
  margin: 0;
  padding: 0;
  background-color: #9f9f9f;
}
</style>

<style scoped>
.page-background {
  min-height: 100vh;
  width: 100%;
  padding: 40px 0;
}

.app-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 32px;
  background-color: #2a2a2a;
  border-radius: 24px;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.15);

  display: flex;
  flex-direction: column;
}

.top-row {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  padding-bottom: 24px;
  margin-bottom: 24px;
  border-bottom: 1px solid #3d3d3d;
}

.controls-column {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: flex-end;
  gap: 20px;
}

.content-body {
  display: flex;
  flex-direction: column;
  gap: 32px;
}

@media (max-width: 1200px) {
  .app-container {
    margin: 0 20px;
  }
}

@media (max-width: 768px) {
  .top-row {
    flex-direction: column;
    gap: 24px;
    align-items: center;
  }

  .controls-column {
    align-items: center;
    width: 100%;
  }
}
</style>