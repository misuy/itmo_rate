<script setup lang="ts">
import RatingCircle from "@/components/RatingCircle.vue";
import ListHeader from "@/components/ListHeader.vue";
import TeacherReviewCard from "@/components/TeacherReviewCard.vue";
import { useStore } from "vuex";
import { useRoute } from "vue-router";

const store = useStore();
const route = useRoute();

const id = Number(route.params.id);
if (id) {
  store.dispatch("getTeacher", id);
  store.dispatch("getTeacherReviews", {id, offset: 0, amount: 10});
} else {
  store.state.error = 404;
}

</script>

<template>
  <div class="home-view">
    <div class="content">
      <div class="page-header">
        <h2 class="prefix-h" @click="$router.back()">
          Преподаватели \ 
        </h2>
        <h2> {{ $store.state.currentTeacher.name }}</h2>
        <div class="h-line" />
      </div>
      <div class="page-content">
        <div class="main-column">
          <div class="subject-info-block">
            <div>
              <div class="subject-info-left">
                <img class="teacher-avatar" :src="$store.state.currentTeacher.avatar">
                <div>
                  <h3>Оценки</h3>
                  <div class="avg-rating-container">
                    <div v-for="cr in $store.state.currentTeacher.criteria" :key="cr.name" class="avg-rating">
                      <RatingCircle :rating="cr.rating" />
                      <span>{{ cr.name }}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="reviews-block">
            <TeacherReviewCard
              v-for="rev in $store.state.teacherReviews" :key="rev.id"
              :date="rev.created" :text="rev.text" :subject="rev.subject" :score="rev.rating"
            />
          </div>
        </div>
        <div class="side-column">
          <div class="classes-block yellow">
            <ListHeader name="Предметы" :count="$store.state.currentTeacher.subjects.length" :font-size="20" />
            <ul>
              <li v-for="subj in $store.state.currentTeacher.subjects" :key="subj"> 
                {{ subj }}
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
  font-size: 20px;
  font-family: Inter,serif;
  font-weight: 600;
  color: var(--text-color);
  margin-bottom: 12px;
  margin-top: 0px;
}

.classes-block {
  margin-top: 25px;
  position: relative;
}
.teachers-block
.classes-block li::marker {
  content: ' ';
}

.classes-block li {
  margin-bottom: 5px;
}

.classes-block li::before {
  content: ' ';
  width: 12px;
  height: 12px;
  border-radius: 15px;
  position: absolute;
  left: 18px;
  margin-top: 3px;
}

.classes-block.yellow li::before {
  border: solid var(--strokes-yellow-color) 1px;
  background: var(--background-yellow-color);
}

@media screen and (max-width: 1300px) {
  .reviews-block {
    columns: 1 auto;
    margin-top: 45px;
  }
}

@media screen and (min-width: 1300px) {
  .reviews-block {
    columns: 2 auto;
    margin-top: 45px;
  }
}

.content {
  /*border: solid red 2px;*/
  /* width: 100%;
  margin-top: 30px;
  margin-left: 3em;
  margin-right: 3em; */
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

.subject-info-left {
  display: flex;
  flex-direction: row;
  gap: 25px;
  margin-top: 25px;
}

.teacher-avatar{
  max-width: 230px;
  max-height: 305px;
  border-radius: 15px;
}

.subject-info-right {
  display: flex;
  flex-direction: column;
  justify-content: start;
  margin-top: 25px;
  gap: 25px;
  align-items: center;
}

.main-column {
  flex-shrink: 10;
}

.avg-rating-container {
  display: flex;
  flex-flow: row wrap;
  gap: 12px 20px;
  max-width: 550px;
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