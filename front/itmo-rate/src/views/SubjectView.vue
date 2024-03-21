<script>
import LeftPanelView from "@/views/LeftPanelView.vue";
import ChipComponent from "@/components/ChipComponent.vue";
import RatingCircle from "@/components/RatingCircle.vue";
import CommonButton from "@/components/CommonButton.vue";

export default {
  name: "SubjectView",
  components: {CommonButton, RatingCircle, ChipComponent, LeftPanelView}
}
</script>

<template>
  <div class="home-view">
    <LeftPanelView />
    <div class="content">
      <div class="page-header">
        <h2 class="prefix-h">Предметы \ </h2>
        <h2>Основы проектной деятельности</h2>
        <div class="h-line"></div>
      </div>
      <div class="page-content">
        <div class="main-column">
          <div class="subject-info-block">
            <div class="subject-info-left">
              <h3>Факультеты</h3>
              <div class="chip-container">
                <ChipComponent text="ПИикт" color="yellow"/>
                <ChipComponent text="ВТ" color="blue"/>
              </div>
              <h3>Оценки</h3>
              <div class="avg-rating-container">
                <div class="avg-rating">
                  <RatingCircle rating="4.5"/>
                  <span>Критерий 1</span>
                </div>
                <div class="avg-rating">
                  <RatingCircle rating="7.5"/>
                  <span>Критерий 2</span>
                </div>
                <div class="avg-rating">
                  <RatingCircle rating="9"/>
                  <span>Критерий 3</span>
                </div>
                <div class="avg-rating">
                  <RatingCircle rating="1.5"/>
                  <span>Критерий 4</span>
                </div>
                <div class="avg-rating">
                  <RatingCircle rating="4.5"/>
                  <span>Критерий 5</span>
                </div>
              </div>
            </div>
            <div class="subject-info-right">
              <RatingCircle rating="6.5" radius="100"/>
              <CommonButton style="font-size: 20px; height: 40px; width: 250px">
                <template #text> Добавить отзыв </template>
                <template #icon>
                  <img src="../assets/icons/PlusIcon.svg" />
                </template>
              </CommonButton>
            </div>
          </div>
          <div class="reviews-block">

          </div>
        </div>
        <div class="side-column">
          <div class="lecturers-block">

          </div>
          <div class="teachers-block">

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

.content {
  /*border: solid red 2px;*/
  width: 80%;
  margin-top: 30px;
  margin-left: auto;
  margin-right: auto;
}

.page-content {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.subject-info-block {
  display: flex;
  flex-direction: row;
  gap: 10px;
}

.subject-info-left {
  max-width: 450px
}
.subject-info-right {
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  gap: 32px;
  align-items: center;
}

.side-column {
  width: 380px;
  background: red;
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

.h-line {
  width: 100%;
  height: 1px;
  background: var(--strokes-color);
  margin-top: 5px;
}

.home-view {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: row;
}
</style>