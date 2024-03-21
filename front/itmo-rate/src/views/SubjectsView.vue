<script setup lang="ts">
import LeftPanelView from './LeftPanelView.vue';
import { RatedListItemInfo } from '@/components/RatedListItem.vue';
import RatedList from '@/components/RatedList.vue';

const testSubjects = [new RatedListItemInfo(6.8, 'Основы проектной деятельности', ['ПИиКТ', 'BT']), new RatedListItemInfo(2.8, 'ТПО', ['ПИиКТ', 'BT', 'Нейротех'])];
</script>

<template>
    <div class="subjects-view">
      <LeftPanelView />
      <div class="subjects-view-content">
        <div class="subjects-list">
          <RatedList :name="'Предметы'" :items="testSubjects" />
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
    margin-top: 50px;
    width: 85%;
  }
</style>
  