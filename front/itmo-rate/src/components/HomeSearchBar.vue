<script setup lang="ts">
import { ref, watch } from 'vue';
</script>

<template>
  <div class="home-search-bar">
    <input class="home-search-bar-input" type="text" placeholder="Преподаватель, предмет..." />
    <img class="magnifier-icon" src="@/assets/icons/MagnifierIcon.svg" />
  </div>
</template>

<style>
.home-search-bar {
  /* width: 100%; */
  display: flex;
  position: relative;
  align-items: center;
}

.home-search-bar-input {
  width: 100%;
  height: 40px;
  border: 1px solid #c3cde1;
  border-radius: 25px;
  padding-left: 20px;
  font-size: 17px;
  color: #656d7c;
}

.magnifier-icon {
  height: 25px;
  width: auto;
  position: absolute;
  right: 12px;
  user-select: none;
  -webkit-user-drag: none;
  background-color: white;
  border-radius: 10px;
}

.home-search-bar-input:focus {
  color: black;
  border: 1px solid #656d7c;
  box-shadow: none;
  outline: none;
}
</style>
