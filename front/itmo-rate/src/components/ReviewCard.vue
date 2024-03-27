<script lang="ts">
import RatingBar from '@/components/RatingBar.vue'
export default {
  name: "ReviewCard",
  components: {RatingBar}
}
</script>

<template>
  <div class="card-body">
    <RatingBar />
    <div class="h-line" style="margin-top: 15px;" />
    <p>
      Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat laboris... 
    </p>
  </div>
</template>

<style scoped>
.card-body {
  background: var(--panel-background-color);
  border-radius: 15px;
  padding: 20px;
  max-width: 360px;
}

</style>