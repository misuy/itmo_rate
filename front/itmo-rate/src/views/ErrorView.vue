<script setup lang="ts">

</script>

<template>
  <div class="content">
    <div class="container">
      <h1> {{ $store.state.error }}</h1>
      <img src="@/assets/star.svg">
    </div>
    <span>Ошибка загрузки страницы! Что-то пошло не так...</span>
  </div>
</template>

<style scoped>

.content {
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: start;
  align-items: center;
  margin-top: 100px;
}

.container {
  position: relative;
  margin-bottom: 30px;
}

h1 {
  font-size: 72px;
  font-weight: 700;
  color: var(--text-color-2);
  z-index: 1;
  padding-top: 10px;
  padding-right: 100px;
}

img {
  position: absolute;
  top: 0;
  z-index: -1;
}

span {
    font-size: 18px;
    font-weight: 500;
    color: var(--text-color);
}

</style>