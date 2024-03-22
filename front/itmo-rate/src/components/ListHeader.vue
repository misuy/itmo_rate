<script setup lang="ts">
const props = defineProps({
    name: String,
    fontSize: {
      type: Number,
      default: 23
    },
    count: {
        type: Number,
        default: -1,
    },
})
</script>

<template>
    <div class="rated-list-header">
        <div class="rated-list-name" v-bind:style="{fontSize: fontSize + 'px'}">
            {{ props.name }}
        </div>
        <div v-if="count != -1" class="rated-list-items-count" v-bind:style="{fontSize: fontSize + 'px'}">
            {{ props.count }}
        </div>
    </div>
    <div class="h-line"></div>
</template>

<style scoped>

.rated-list-header {
  display: flex;
  align-items: center;
  font-weight: 600;
}

.rated-list-name {
  color: black;
  margin-right: 7px;
}

.rated-list-items-count {
  color: #D7DFEC;
}
</style>