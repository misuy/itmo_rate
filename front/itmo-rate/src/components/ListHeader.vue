<script setup lang="ts">
const props = defineProps({
    name: String,
    count: {
        type: Number,
        default: -1,
    },
})
</script>

<template>
    <div class="rated-list-header">
        <div class="rated-list-name">
            {{ props.name }}
        </div>
        <div v-if="count == -1" class="rated-list-items-count">
            {{ props.count }}
        </div>
    </div>
</template>

<style scoped>
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
</style>