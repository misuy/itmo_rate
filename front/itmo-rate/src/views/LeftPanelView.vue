<script setup lang="ts">
import NavBox from '@/components/NavBox.vue'
import UserBox from '@/components/UserBox.vue'
import { watch } from 'vue';
import { useRoute } from 'vue-router';
import { useStore } from 'vuex';

const route = useRoute();
const store = useStore();
watch(
  () => route.fullPath,
  () => { store.state.error = NaN; }
);

</script>

<template>
  <div class="left-panel">
    <img class="itmo-logo" src="@/assets/itmo_logo.svg">

    <div class="nav-box-holder">
      <NavBox />
    </div>

    <div class="user-box-holder">
      <UserBox first-name="123" second-name="321" />
    </div>
  </div>
</template>

<style>
.left-panel {
  background: var(--panel-background-color);
  display: flex;
  align-items: center;
  flex-direction: column;
  position: fixed;
  z-index: 1;
  height: 100%;
  min-width: 220px;
}

.itmo-logo {
  height: 75px;
  margin-top: 25px;
  margin-bottom: 20px;
}

.nav-box-holder {
  width: 70%;
  margin-bottom: 40px;
}

.user-box-holder {
  width: 77%;
  margin-top: auto;
  margin-bottom: 30px;
}
</style>
