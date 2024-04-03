<script setup lang="ts">
import type { TeacherPreview } from '@/api/teachers';
import ListHeader from './ListHeader.vue';
import RatedListItem from './RatedListItem.vue';
import type { PropType } from 'vue';

interface Info {
  score: number;
  name: string;
  tags: string[];
}

defineProps({
    name: String,
    items: Object as PropType<Info[]>,
})
</script>

<template>
  <div class="rated-list">
    <ListHeader :name="name" :count="items?.length" />
    <div class="rated-list-items">
      <RatedListItem v-for="item in items" :info="item" :key="item.name" />
    </div>
  </div>
</template>

<style>
.rated-list {
    width: 100%;
    display: flex;
    flex-direction: column;
}

.rated-list-items {
    display: flex;
    flex-direction: column;
}
</style>