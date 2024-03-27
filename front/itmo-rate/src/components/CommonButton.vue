<script lang="ts">
export default {
  name: "CommonButton",
  props: {
    text: String,
    icon: {
      type: Boolean,
      default: false
    }
  }
}
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
  background: var(--btn-color);
  transition: background .3s;
  cursor: pointer;
  width: fit-content;
}

.button:hover {
  background: var(--btn-hover-color);
  transition: background .3s;
}

.button span {
  color: #FFFFFF;
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