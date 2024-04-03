<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent ({
  name: "OutlineButton",
  props: {
    text: String,
    icon: {
      type: Boolean,
      default: false
    }
  }
})
</script>

<template>
  <div class="button">
    <span>
      {{ text }}
    </span>
    <div v-if="icon" class="icon">
      <slot name="icon" />
    </div>
  </div>
</template>

<style scoped>
.button {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  font-weight: 500;
  border-radius: 12px;
  gap: 10px;
  padding: 12px 20px;
  transition: background .3s; 
  transition: color .3s; 
  border: solid var(--strokes-color) 1px;
  color: var(--text-color-2);
  cursor: pointer;
  width: fit-content;
}

.button:hover {
  background: var(--text-color-2);
  border: solid var(--text-color-2) 1px;
  color: white;
  transition: background .3s;
}

.button span {
  font-weight: 500;
  max-lines: 1;
}

.button:hover *{
  opacity: 0.9;
  transition: opacity .3s;
}

.icon {
  max-width: 18px;
  max-height: 18px;
}

</style>