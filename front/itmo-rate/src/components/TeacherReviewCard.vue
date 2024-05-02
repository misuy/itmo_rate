<script lang="ts">
import { defineComponent } from 'vue';
import CommonButton from '@/components/CommonButton.vue';

export default defineComponent({
  name: "TeacherReviewCard",
  components: {CommonButton},
  props: {
    id: {
      type: Number,
      required: true
    },
    text: {
      type: String,
      default: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat laboris, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nostrud exercitation"
    },
    subject: {
      type: String,
      default: "<subject>"
    },
    date: {
      type: String,
      default: "<date>"
    },
    score: {
      type: Number,
      default: 0
    }
  },
  computed: {
    truncText() : string {
      return this.text.length > 373 ? this.text.substring(0, 373) + "..." : this.text;
    },
    computeColor() {
      if (this.score <= 2.5)
        return "#FF0000";
      else if (this.score <= 5)
        return "#E3B064";
      else if (this.score <= 7.5)
        return "#ABE364";
      else
        return "#85E364";
    },
    getDate() : String {
      return new Date(this.date).toLocaleString()
    }
  }
})
</script>

<template>
  <div class="card-body">
    <div class="top-bar">
      <div class="rating-container" :style="{background: computeColor}">
        {{ score.toFixed(1) }}
      </div>
      <div>
        <h4>{{ subject }}</h4>
        <span> {{ getDate }}</span>
      </div>
    </div>
    <div class="h-line" style="margin-top: 15px;" />
    <p>
      {{ truncText }}
    </p>
    <div class="under-block">
      <CommonButton
        text="Читать" style="width: 80px;" 
        @click="() => { $store.dispatch('getReview', id); }"
      />
      <span>
        Анонимно
      </span>
    </div>
  </div>
</template>

<style scoped>

h4 {
  font-size: 18px;
  font-weight: 500;
  color: var(--text-color);
  margin: 0;
  /* margin-bottom: 5px; */
}

span {
  font-size: 16px;
  color: var(--text-color-2);
}

.top-bar {
  display: flex;
  flex-direction: row;
  gap: 15px;
}

.rating-container {
  width: 57px;
  height: 63px;
  border-radius: 10px 0px 0px 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 22px;
  color: white;
}

.under-block {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: end;
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