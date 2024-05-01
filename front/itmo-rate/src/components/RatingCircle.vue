<script lang="ts">
export default {
  name: "RatingCircle",
  props: {
    rating: {
      type: Number,
      required: true
    },
    radius: {
      type: Number,
      default: 24,
    }
  },
  computed: {
    computeColor() {
      if (this.rating <= 2.5)
        return "#FF0000";
      else if (this.rating <= 5)
        return "#E3B064";
      else if (this.rating <= 7.5)
        return "#ABE364";
      else
        return "#85E364";
    },
    computeSize() {
      return (this.radius * 2).toString() + "px";
    }
  }
}
</script>

<template>
  <div :style="{background: computeColor, width: computeSize, height: computeSize, borderRadius: computeSize}">
    <span> {{ rating.toFixed(1) }} </span>
  </div>
</template>

<style scoped>
div{
  display: flex;
  align-items: center;
  justify-content: center;
}
span {
  color: #FFFFFF;
  font-weight: 600;
  font-size: 18px;
  text-align: center;
}
</style>