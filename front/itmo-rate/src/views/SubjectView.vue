<script lang="ts">
import ChipComponent from "@/components/ChipComponent.vue";
import RatingCircle from "@/components/RatingCircle.vue";
import CommonButton from "@/components/CommonButton.vue";
import ListHeader from "@/components/ListHeader.vue";
import ReviewCard from "@/components/ReviewCard.vue";

export default {
  name: "SubjectView",
  components: {ListHeader, CommonButton, RatingCircle, ChipComponent, ReviewCard},
  data() {
    return {
      lecturers: ["Клименков Сергей Викторович", "Соснов Николай Федорович"],
      teachers: ["Тимофеев Тихон Александрович", "Колпакова Екатерина Александровна", "Рябов Лука Макарович"]
    }
  }
}
</script>

<template>
  <div class="home-view">
    <div class="content">
      <div class="page-header">
        <h2 class="prefix-h">
          Предметы \ 
        </h2>
        <h2>Основы проектной деятельности</h2>
        <div class="h-line" />
      </div>
      <div class="page-content">
        <div class="main-column">
          <div class="subject-info-block">
            <div class="subject-info-left">
              <h3>Факультеты</h3>
              <div class="chip-container">
                <ChipComponent text="ПИикт" color="yellow" />
                <ChipComponent text="ВТ" color="blue" />
              </div>
              <h3>Оценки</h3>
              <div class="avg-rating-container">
                <div class="avg-rating">
                  <RatingCircle :rating="4.5" />
                  <span>Критерий 1</span>
                </div>
                <div class="avg-rating">
                  <RatingCircle :rating="7.5" />
                  <span>Критерий 2</span>
                </div>
                <div class="avg-rating">
                  <RatingCircle :rating="9" />
                  <span>Критерий 3</span>
                </div>
                <div class="avg-rating">
                  <RatingCircle :rating="1.5" />
                  <span>Критерий 4</span>
                </div>
                <div class="avg-rating">
                  <RatingCircle :rating="4.5" />
                  <span>Критерий 5</span>
                </div>
              </div>
            </div>
            <div class="subject-info-right">
              <RatingCircle :rating="6.5" :radius="100" />
              <CommonButton style="font-size: 20px; height: 40px; width: 260px">
                <template #text>
                  Добавить отзыв 
                </template>
                <template #icon>
                  <img src="../assets/icons/PlusIcon.svg">
                </template>
              </CommonButton>
            </div>
          </div>
          <div class="reviews-block">
            <ReviewCard />
            <ReviewCard />
            <ReviewCard />
            <ReviewCard />
          </div>
        </div>
        <div class="side-column">
          <div class="teachers-block yellow">
            <ListHeader name="Лекторы" :count="lecturers.length" :font-size="20" />
            <ul>
              <li v-for="lecturer in lecturers" :key="lecturer"> 
                {{ lecturer }}
              </li>
            </ul>
          </div>
          <div class="teachers-block blue">
            <ListHeader name="Преподаватели" :count="teachers.length" :font-size="20" />
            <ul>
              <li v-for="teacher in teachers" :key="teacher">
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
  font-size: 28px;
  font-family: Inter,serif;
  font-weight: 600;
  color: var(--text-color);
}

.page-header h2 {
  display: inline;
}

h2.prefix-h {
  color: var(--strokes-color);
}

h3 {
  font-size: 24px;
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
  display: flex;
  flex-wrap: wrap;
  column-count: auto;
  gap: 25px;
  margin-top: 45px;
  /* justify-content: center; */
}

.content {
  /*border: solid red 2px;*/
  width: 100%;
  margin-top: 30px;
  margin-left: 3em;
  margin-right: 3em;
}

.page-content {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  gap: 3em;
}

.subject-info-block {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  /* justify-content: center; */
  gap: 10px;
}

.subject-info-left {
  max-width: 450px
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
  flex-wrap: wrap;
  gap: 10px 25px;
  max-width: 500px;
}

.avg-rating {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 10px;
  font-size: 18px;
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