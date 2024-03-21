<script setup lang="ts">
import { defineProps } from 'vue';
import router from '@/router';

const props = defineProps({
  routeTo: String,
})
</script>

<template>
  <div @click="router.push(routeTo!)" class="nav-item">
    <div class="nav-item-icon">
      <slot name="icon"></slot>
    </div>

    <div class="nav-item-name">
      <slot name="name" class="nav-item-name"></slot>
    </div>
  </div>
</template>

<style>
.nav-item {
  display: flex;
  align-items: center;
  height: 25px;
  width: 100%;
  cursor: pointer;
}

.nav-item-icon {
  color: #656d7c;
  vertical-align: middle;
  height: 100%;
  width: auto;
}

.nav-icon {
  height: 28px;
  width: 28px
}

.nav-item-name {
  color: #656d7c;
  margin-left: 15px;
  /* vertical-align: middle; */
  font-size: 15px;
  user-select: none;
}

.nav-item:hover * {
  opacity: 0.95;
}
</style>
