<script setup lang="ts">
import LeftPanelView from '@/views/LeftPanelView.vue';
import ErrorView from '@/views/ErrorView.vue';
import { ModalState } from '@/store';
import ReviewOpen from '@/components/ReviewOpen.vue'

</script>

<template>
  <div class="main-view">
    <LeftPanelView />
    <div class="main-content">
      <RouterView v-if="!$store.state.error" />
      <ErrorView v-else />
    </div> 
  </div>
  <div class="modal" v-if="$store.state.modalWindowState != ModalState.NOTHING">
    <ReviewOpen v-if="$store.state.modalWindowState == ModalState.REVIEW_READ_MODE" />
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

.modal {
  background: rgba(10, 12, 19, 0.53);
  position: fixed;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  z-index: 100;
  display: flex;
  align-items: center;
  justify-content: center;
}

</style>