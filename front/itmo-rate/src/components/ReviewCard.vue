<script lang="ts">
import RatingBar from '@/components/RatingBar.vue'
import { defineComponent } from 'vue';
import CommonButton from '@/components/CommonButton.vue';

export default defineComponent({
  name: "ReviewCard",
  components: {RatingBar, CommonButton},
  props: {
    text: {
      type: String,
      default: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat laboris, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nostrud exercitation"
    }
  },
  computed: {
    truncText() : string {
      return this.text.length > 375 ? this.text.substring(0, 375) + "..." : this.text;
    }
  }
})
</script>

<template>
  <div class="card-body">
    <RatingBar />
    <div class="h-line" style="margin-top: 15px;" />
    <p>
      {{ truncText }}
    </p>
    <div class="under-block">
      <CommonButton text="Читать" style="width: 80px;" />
      <span>
        12.02.2003<br>
        Анонимно
      </span>
    </div>
  </div>
</template>

<style scoped>

.under-block {
  display: flex;
  flex-direction: row;
  justify-content: start;
  align-items: center;
  gap: 15px;
}

.under-block span {
  font-size: 0.9em;
  color: var(--text-color-2)
}

.card-body {
  background: var(--panel-background-color);
  border-radius: 15px;
  padding: 20px;
  margin-bottom: 25px;
}

</style>