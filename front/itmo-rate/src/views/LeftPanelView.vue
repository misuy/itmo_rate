<script setup lang="ts">
import NavBox from '@/components/NavBox.vue'
import UserBox from '@/components/UserBox.vue'
</script>

<template>
  <div class="left-panel">
    <img class="itmo-logo" src="@/assets/itmo_logo.svg">

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
