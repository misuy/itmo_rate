<script setup lang="ts">
import LeftPanelView from '@/views/LeftPanelView.vue';
</script>

<template>
  <div class="main-view">
    <LeftPanelView />
    <div class="main-content">
      <RouterView />
    </div>
  </div>
</template>

<style scoped>
.main-view {
    width: 100%;
    height: 100%;
    display: flex;
}

.main-content {
  width: 100%;
  height: 100%;
  padding-left: 220px;
  margin-top: 30px;
  margin-left: 3em;
  margin-right: 3em;
}
</style>