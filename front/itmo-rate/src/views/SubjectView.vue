<script setup lang="ts">
import ChipComponent from "@/components/ChipComponent.vue";
import RatingCircle from "@/components/RatingCircle.vue";
import RatingAndReviewCircle from '@/components/RatingAndReviewCircle.vue'
import CommonButton from "@/components/CommonButton.vue";
import ListHeader from "@/components/ListHeader.vue";
import ReviewCard from "@/components/ReviewCard.vue";
import { useStore } from "vuex";
import { useRoute } from "vue-router";

const store = useStore();
const route = useRoute();

const id = Number(route.params.id);
if (id) {
  store.dispatch("getSubject", id);
  store.dispatch("getSubjectReviews", {id, offset: 0, amount: 10});
} else {
  store.state.error = 404;
}

</script>

<template>
  <div class="home-view">
    <div class="content">
      <div class="page-header">
        <h2 class="prefix-h" @click="$router.back()">
          Предметы \ 
        </h2>
        <h2>{{ $store.state.currentSubject.name }}</h2>
        <div class="h-line" />
      </div>
      <div class="page-content">
        <div class="main-column">
          <div class="subject-info-block">
            <div class="subject-info-left">
              <h3>Факультеты</h3>
              <div class="chip-container">
                <ChipComponent 
                  v-for="(tag, idx) in $store.state.currentSubject.faculties" :key="tag" 
                  :text="tag" 
                  :color="idx % 2 == 0 ? 'yellow' : 'blue'"
                />
              </div>
              <h3>Оценки</h3>
              <div class="avg-rating-container">
                <div v-for="cr in $store.state.currentSubject.criteria" :key="cr.name" class="avg-rating">
                  <RatingCircle :rating="cr.rating" />
                  <span>{{ cr.name }}</span>
                </div>
              </div>
            </div>
            <div class="subject-info-right">
              <RatingAndReviewCircle :reviews="$store.state.currentSubject.reviews_count" :rating="$store.state.currentSubject.avg_rating" style="width: 200px; height: 200px;" />
              <CommonButton text="Добавить отзыв" :icon="true" style="font-size: 18px; height: 35px; width: 260px;">
                <template #icon>
                  <img src="../assets/icons/PlusIcon.svg">
                </template>
              </CommonButton>
            </div>
          </div>
          <div class="reviews-block">
            <ReviewCard
              v-for="card in $store.state.subjectReviews" :key="card.id"
              :date="card.created" :text="card.text" :subject="card.subject" :rating="card.rating"
              class="review-card"
            />
          </div>
        </div>
        <div class="side-column">
          <div class="teachers-block yellow">
            <ListHeader name="Лекторы" :count="$store.state.currentSubject.lecturers.length" :font-size="20" />
            <ul>
              <li v-for="lecturer in $store.state.currentSubject.lecturers" :key="lecturer"> 
                {{ lecturer }}
              </li>
            </ul>
          </div>
          <div class="teachers-block blue">
            <ListHeader name="Преподаватели" :count="$store.state.currentSubject.length" :font-size="20" />
            <ul>
              <li v-for="teacher in $store.state.currentSubject.teachers" :key="teacher">
                {{ teacher }}
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

h2 {
  font-size: 24px;
  font-family: Inter,serif;
  font-weight: 600;
  color: var(--text-color);
}

.page-header h2 {
  display: inline;
}

h2.prefix-h {
  color: var(--strokes-color);
  cursor: pointer;
}

h3 {
  font-size: 22px;
  font-family: Inter,serif;
  font-weight: 600;
  color: var(--text-color);
  margin-bottom: 12px;
  margin-top: 25px;
}

.teachers-block {
  margin-top: 25px;
  position: relative;
}

.teachers-block li::marker {
  content: ' ';
}

.teachers-block li::before {
  content: ' ';
  width: 12px;
  height: 12px;
  border-radius: 15px;
  position: absolute;
  left: 15px;
  margin-top: 3px;
}

.teachers-block.yellow li::before {
  border: solid var(--strokes-yellow-color) 1px;
  background: var(--background-yellow-color);
}

.teachers-block.blue li::before {
  border: solid var(--strokes-blue-color) 1px;
  background: var(--background-blue-color);
}

.reviews-block {
  margin-top: 35px;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  flex-wrap: wrap;
}
@media screen and (min-width: 1300px) {
  .review-card {
    width: 43%;
  }
}

@media screen and (max-width: 1300px) {
  .review-card {
    width: 100%;
  }
}


.content {
  /*border: solid red 2px;*/
  width: 100%;
}

.page-content {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  gap: 80px;
}

.subject-info-block {
  display: flex;
  flex-direction: row;
  gap: 10px;
  justify-content: space-between;
  /* max-width: 760px */
  width: auto;
}

.subject-info-right {
  display: flex;
  flex-direction: column;
  justify-content: start;
  margin-top: 25px;
  gap: 25px;
  align-items: center;
}

.side-column {
  /* width: 380px; */
  /*background: red;*/
}
.main-column {
  flex-shrink: 10;
}

.avg-rating-container {
  display: flex;
  flex-flow: row wrap;
  gap: 10px 20px;
  /* max-width: 350px; */
}

.avg-rating {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 8px;
  font-size: 16px;
  font-weight: 500;
}

.chip-container {
  display: flex;
  gap: 10px;
}

.home-view {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: row;
}
</style>