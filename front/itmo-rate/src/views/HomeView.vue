<script setup lang="ts">
import LeftPanelView from '@/views/LeftPanelView.vue'
import HomeSearchBar from '@/components/HomeSearchBar.vue'
import { RatedListItemInfo } from '@/components/RatedListItem.vue';
import RatedList from '@/components/RatedList.vue';
import { computed, ref } from 'vue';

const testSubjects = [new RatedListItemInfo(6.8, 'Основы проектной деятельности', ['ПИиКТ', 'BT']), new RatedListItemInfo(2.8, 'ТПО', ['ПИиКТ', 'BT', 'Нейротех'])];
const testTeachers = [new RatedListItemInfo(4.4, 'Соснов Николай Федорович', ['Методы криптографии', 'Компьютерные сети', 'ТПО']), new RatedListItemInfo(7.2, 'Соснов Семен Федорович', ['Методы криптографии', 'Компьютерные сети', 'ТПО'])]

const substr = ref('')
const relatedSubjects = computed(() => testSubjects.filter((item) => item.name.toLowerCase().includes(substr.value.toLowerCase())))
const relatedTeachers = computed(() => testTeachers.filter((item) => item.name.toLowerCase().includes(substr.value.toLowerCase())))

</script>

<template>
  <div class="home-view">
    <LeftPanelView />
    <div class="home-view-content">
      <img class="itmo-rate-logo" src="@/assets/itmo_rate_logo.svg" />
      <div class="home-search-bar-holder">
        <HomeSearchBar @input-changed="(input) => substr=input"/>
      </div>
      <div class="subjects-list">
        <RatedList :name="'Предметы'" :items="relatedSubjects" />
      </div>
      <div class="teachers-list">
        <RatedList :name="'Преподаватели'" :items="relatedTeachers" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.home-view {
  width: 100%;
  height: 100%;
  display: flex;
}

.home-view-content {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.itmo-rate-logo {
  width: auto;
  height: 80px;
  margin-top: 88px;
}

.home-search-bar-holder {
  margin-top: 15px;
  width: 45%;
}

.subjects-list {
  margin-top: 30px;
  width: 75%;
}

.teachers-list {
  margin-top: 30px;
  width: 75%;
}
</style>
