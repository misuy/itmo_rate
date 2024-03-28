<script setup lang="ts">
import LeftPanelView from '@/views/LeftPanelView.vue';
import TeacherView from '@/views/TeacherView.vue';
</script>

<template>
  <div class="main-view">
    <LeftPanelView />
    <div class="main-content">
      <!-- <RouterView /> -->
      <TeacherView />
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
    margin-left: 220px;
}
</style>