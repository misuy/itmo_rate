<script setup lang="ts">
import HomeSearchBar from '@/components/HomeSearchBar.vue'
import RatedList from '@/components/RatedList.vue';
import { useStore } from 'vuex';
import { SearchState } from '@/store';

const store = useStore();
store.commit("gotTeachers", []);
store.commit("gotSubjects", []);
let readyForSearch = true;
let lastSearched = "<>"

function search(input: string) {
  lastSearched = input;
  if (input == "") {
      store.commit("setSearchState", SearchState.IDLE);
      return;
  }
  if (!readyForSearch)
    return;
  store.dispatch("search", input);
  console.log(`Search request: ${input}`);
  readyForSearch = false;
  setTimeout(() => { 
    readyForSearch = true;
    if (input != lastSearched && lastSearched != "") {
      console.log(`Search request: ${lastSearched}`);
      store.dispatch("search", lastSearched);
    } 
  }, 1000)
}

</script>

<template>
  <div class="home-view">
    <div class="home-view-content">
      <img class="itmo-rate-logo" src="@/assets/itmo_rate_logo.svg">
      <div class="home-search-bar-holder">
        <HomeSearchBar @input-changed="search" />
      </div>
      <div class="sub-list" v-if="$store.state.subjects.length > 0 && $store.state.searchState == SearchState.FOUND">
        <RatedList :name="'Предметы'" :items="$store.state.subjects" path-prefix="subject/" />
      </div>
      <div class="sub-list" v-if="$store.state.teachers.length > 0 && $store.state.searchState == SearchState.FOUND">
        <RatedList :name="'Преподаватели'" :items="$store.state.teachers" path-prefix="teacher/" />
      </div>
      <div class="sub-list" v-if="$store.state.searchState == SearchState.FETCHING">
        Загрузка...
      </div>
      <div class="sub-list" v-else-if="$store.state.searchState == SearchState.NOTHING_FOUND">
        Ничего не найдено!
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
  margin-top: 48px;
}

.home-search-bar-holder {
  margin-top: 15px;
  width: 45%;
}

.sub-list {
  margin-top: 30px;
  width: 75%;
  display: flex;
  justify-content: center;
}

</style>
