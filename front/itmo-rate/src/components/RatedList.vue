<script setup lang="ts">
import { reactive } from 'vue';
import ListHeader from './ListHeader.vue';
import RatedListItem, { RatedListItemInfo } from './RatedListItem.vue';

defineProps({
    name: String,
    items: Array<RatedListItemInfo>,
})
</script>

<template>
    <div class="rated-list">
        <ListHeader :name="name" :count="items?.length" />
        <hr size="1" class="rated-list-divider" />
        <div class="rated-list-items">
            <RatedListItem v-for="item in items" :info="item" />
        </div>
    </div>
</template>

<style>
.rated-list {
    width: 100%;
    display: flex;
    flex-direction: column;
}

.rated-list-divider {
    color: var(--strokes-color);
    width: 100%;
    margin: 0;
}

.rated-list-items {
    display: flex;
    flex-direction: column;
}
</style>