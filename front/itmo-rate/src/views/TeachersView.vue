<script setup lang="ts">
import LeftPanelView from './LeftPanelView.vue';
import { RatedListItemInfo } from '@/components/RatedListItem.vue';
import RatedList from '@/components/RatedList.vue';

const testTeachers = [new RatedListItemInfo(4.4, 'Соснов Николай Федорович', ['Методы криптографии', 'Компьютерные сети', 'ТПО']), new RatedListItemInfo(7.2, 'Соснов Семен Федорович', ['Методы криптографии', 'Компьютерные сети', 'ТПО'])]
</script>

<template>
  <div class="teachers-view">
    <div class="teachers-view-content">
      <div class="teachers-list">
        <RatedList :name="'Преподаватели'" :items="testTeachers" />
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
