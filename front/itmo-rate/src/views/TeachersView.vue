<script setup lang="ts">
import RatedList from '@/components/RatedList.vue';
import { useStore } from 'vuex';

const store = useStore();
store.dispatch("gotTeachers", {
  offset: 0,
  amount: 10
})

</script>

<template>
  <div class="teachers-view">
    <div class="teachers-view-content">
      <div class="teachers-list">
        <RatedList :name="'Преподаватели'" :items="$store.state.teachers" />
      </div>
    </div>
  </div>
</template>
  
<style scoped>
  .teachers-view {
    width: 100%;
    height: 100%;
    display: flex;
  }
  
  .teachers-view-content {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  
  .teachers-list {
    margin-top: 50px;
    width: 85%;
  }
</style>
