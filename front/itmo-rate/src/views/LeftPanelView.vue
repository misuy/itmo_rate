<script setup lang="ts">
import NavBox from '@/components/NavBox.vue'
import UserBox from '@/components/UserBox.vue'
</script>

<template>
  <div class="left-panel">
    <img class="itmo-logo" src="@/assets/itmo_logo.svg" />

    <div class="nav-box-holder">
      <NavBox />
    </div>

    <div class="user-box-holder">
      <UserBox />
    </div>
  </div>
</template>

<style>
.left-panel {
  background-color: #eff1f6;
  display: flex;
  align-items: center;
  flex-direction: column;
  height: 100%;
  min-width: 220px;
}

.itmo-logo {
  height: 75px;
  margin-top: 25px;
  margin-bottom: 20px;
}

.user-box-holder {
  position: sticky;
  width: 77%;
  top: 83%;
}

.nav-box-holder {
    width: 70%;
}
</style>
