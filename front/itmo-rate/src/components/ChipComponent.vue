<script>

export default {
  name: "ChipComponent",
  props: {
    text: String,
    color: {
      type: String,
      default: "yellow",
      validator: (value) => ["yellow", "blue"].includes(value)
    }
  },
  computed: {
    yellowStyle() {
      return this.color === 'yellow'
    },
    blueStyle() {
      return this.color !== 'yellow'
    }
  }
}
</script>

<template>
  <div class="chip-body" :class="{'yellow-style': yellowStyle, 'blue-style': blueStyle}">
    <span>{{text}}</span>
  </div>
</template>

<style scoped>
.yellow-style {
  background: var(--background-yellow-color);
  border: solid 1px var(--strokes-yellow-color);
  color: var(--text-yellow-color);
}

.blue-style {
  background: var(--background-blue-color);
  border: solid 1px var(--strokes-blue-color);
  color: var(--text-blue-color);
}

.chip-body {
  height: 26px;
  border-radius: 14px;
  padding-left: 8px;
  padding-right: 8px;
  display: flex;
  align-items: center;
}

.chip-body span {
  display: inline;
  font-size: 16px;
  text-align: center;
}
</style>