<script setup lang="ts">
import { reactive } from 'vue';
import RatedListItem, { RatedListItemInfo } from './RatedListItem.vue';

const props = defineProps({
    name: String,
    items: Array<RatedListItemInfo>,
})
</script>

<template>
    <div class="rated-list">
        <div class="rated-list-header">
            <div class="rated-list-name">
                {{ props.name }}
            </div>
            <div class="rated-list-items-count">
                {{ props.items?.length }}
            </div>
        </div>
        <hr size="1" class="rated-list-divider" />
        <div class="rated-list-items">
            <RatedListItem v-for="item in props.items" :info="item" />
        </div>
    </div>
</template>

<style>
.rated-list {
    width: 100%;
    display: flex;
    flex-direction: column;
}

.rated-list-header {
    display: flex;
    align-items: center;
    font-size: 23px;
    font-weight: 600;
    margin-bottom: 3px;
}

.rated-list-name {
    color: black;
    margin-right: 7px;
}

.rated-list-items-count {
    color: #D7DFEC;
}

.rated-list-divider {
    color: #c3cde1;
    width: 100%;
    margin: 0;
}

.rated-list-items {
    display: flex;
    flex-direction: column;
}
</style>