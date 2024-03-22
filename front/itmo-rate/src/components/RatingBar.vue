<script lang="ts">

export default {
  name: "RatingBar",
  props: {
    score: {
      type: Array<Number>,
      default: [0, 0, 0, 0, 0],
      validator: (value) => value.length == 5
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
    }
  }
}
</script>

<template>
  <div class="container">
    <div v-for="s in score" class="rating-rect">
      <div class="rating" :style="{background: computeColor}">
        <spa>{{ s }}</spa>
      </div>
      <div class="char">A</div>
    </div>
  </div>
</template>

<style scoped>

.container {
  display: flex;
  flex-direction: row;
  gap: 15px;
  font-size: 18px;
  font-weight: 700;
}

.rating-rect {
  width: 62px;
  height: 72px;
}

.rating {
  width: 100%;
  height: 42px;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px 8px 0px 0px;
}

.char {
  width: 100%;
  padding-top: 10px;
  text-align: center;
  color: var(--text-color);
}

</style>