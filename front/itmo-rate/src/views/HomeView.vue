<script setup lang="ts">
import LeftPanelView from '@/views/LeftPanelView.vue'
import HomeSearchBar from '@/components/HomeSearchBar.vue'
import Subject from '@/components/Subject.vue';
import { SubjectInfo, SubjectOwner } from '@/components/Subject.vue';

const testSubjectInfo = new SubjectInfo(6.8, 'Основы проектной деятельности', new SubjectOwner('ПИиКТ', 'BT'));
</script>

<template>
  <div class="home-view">
    <LeftPanelView />
    <div class="home-view-content">
      <img class="itmo-rate-logo" src="@/assets/itmo_rate_logo.svg" />
      <div class="home-search-bar-holder">
        <HomeSearchBar />
      </div>
      <Subject :info="testSubjectInfo"/>
    </div>
  </div>
</template>

<style>
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
</style>
