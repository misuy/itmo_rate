<script lang="ts">
export default {
  name: "RatingAndReviewCircle",
  props: {
    rating: {
      type: Number,
      default: 0
    },
    reviews: {
      type: Number,
      default: 0,
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
  <div class="container" :style="{background: computeColor}">
    <span class="rating"> {{ rating.toFixed(2) }} </span>
    <span class="reviews"> {{ reviews + " отзывов" }} </span>
  </div>
</template>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  font-size: 36px;
  align-items: center;
  justify-content: center;
  border-radius: 100%;
}
.rating {
  color: #FFFFFF;
  font-weight: 600;
  text-align: center;
}

.reviews {
  color: #FFFFFF;
  font-weight: 500;
  opacity: 0.8;
  font-size: 0.5em;
  text-align: center;
}
</style>