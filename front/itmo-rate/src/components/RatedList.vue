<script setup lang="ts">
import { useStore } from 'vuex';
import ListHeader from './ListHeader.vue';
import RatedListItem from './RatedListItem.vue';
import type { PropType } from 'vue';

interface Info {
  id: number;
  score: number;
  name: string;
  tags: string[];
}

defineProps({
    name: String,
    pathPrefix: {
      type: String,
      default: "/"
    },
    items: Object as PropType<Info[]>,
})
</script>

<template>
  <div class="rated-list">
    <ListHeader :name="name" :count="items?.length" />
    <div class="rated-list-items">
      <RatedListItem v-for="item in items" :info="item" :key="item.name" :path-prefix="pathPrefix" />
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
    margin-top: 10px;
}
</style>