<script setup lang="ts">
import RatedList from '@/components/RatedList.vue';
import { useStore } from 'vuex';

const store = useStore();
store.dispatch("getSubjects", {
  offset: 0,
  amount: 10
});
console.log(store.state.subjects);
</script>

<template>
  <div class="subjects-view">
    <div class="subjects-view-content">
      <div class="subjects-list">
        <RatedList :name="'Предметы'" :items="$store.state.subjects" path-prefix="subject/" />
      </div>
    </div>
  </div>
</template>
  
<style scoped>
  .subjects-view {
    width: 100%;
    height: 100%;
    display: flex;
  }
  
  .subjects-view-content {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  
  .subjects-list {
    width: 100%;
  }
</style>
  