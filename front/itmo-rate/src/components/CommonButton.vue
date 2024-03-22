<template>
<div class="button">
  <span>
    <slot name="text"></slot>
  </span>
  <div class="icon">
    <slot name="icon"></slot>
  </div>
</div>
</template>

<script>
export default {
  name: "CommonButton"
}
</script>

<style scoped>
.button {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  font-weight: 500;
  border-radius: 10px;
  gap: 10px;
  padding: 10px 20px;
  background: var(--text-color-2);
}

.button span {
  color: #FFFFFF;
  font-weight: 500;
  max-lines: 1;
}

.button :hover {
  opacity: 0.9;
}

.icon {
  max-width: 18px;
  max-height: 18px;
}

</style>