<script lang="ts">

export default {
  name: "RatingBar",
  props: {
    score: {
      type: Array<number>,
      default: [0, 0, 0, 0, 0],
      validator: (value: Array<Number>) => value.length == 5
    },
  },
  methods: {
    computeColor(rating: number) {
      if (rating <= 2.5)
        return "#FF0000";
      else if (rating <= 5)
        return "#E3B064";
      else if (rating <= 7.5)
        return "#ABE364";
      else
        return "#85E364";
    }
  }
}
</script>

<template>
  <div class="container">
    <div v-for="(s, i) in score" :key="i" class="rating-rect">
      <div class="rating" :style="{background: computeColor(s)}">
        <span>{{ s }}</span>
      </div>
      <div class="char">
        {{ "K" + (i + 1) }}
      </div>
    </div>
  </div>
</template>

<style scoped>

.container {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  gap: 5px;
  justify-content: space-between;
  font-size: 18px;
  font-weight: 700;
}

.rating-rect {
  width: 62px;
  height: 100%;
}

.rating {
  width: 100%;
  height: 50%;
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
  height: 50%;
  /* color: var(--text-color-2); */
  font-weight: 600;
}

</style>