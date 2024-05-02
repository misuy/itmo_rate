<script setup lang="ts">
import { ModalState } from '@/store';
import CommonButton from './CommonButton.vue';
import RatingBar from './RatingBar.vue';
import { useStore } from 'vuex';
import type { Criterion } from '@/api';

const store = useStore();
const getRating = () => {
  const result: number[] = [];
  store.state?.openReview?.criteria?.forEach((it: Criterion) => {
    result.push(it.rating);
  });
  return result;
}

const getDate = () => {
  console.log("redraw");
  if (store.state?.openReview?.created) 
    return new Date(store.state?.openReview?.created).toLocaleString();
  return "";
}

</script>

<template>
  <div class="modal-window">
    <div class="top-block">
      <div class="teachers-block">
        <h4> Преподаватель: </h4>
        <span> {{ $store.state?.openReview?.lecturers[0] }}</span>
        <h4 style="margin-top: 5px;">
          Практик:
        </h4>
        <span> {{ $store.state?.openReview?.teachers[0] }} </span>
      </div>
      <RatingBar style="height: 92px;" :score="getRating()" />
    </div>
    <div class="h-line" />
    <div class="text">
      <p>
        {{ $store.state?.openReview?.text }}
      </p>
    </div>
    <div class="bottom-block">
      <CommonButton
        text="Закрыть" style="width: 90px;" 
        @click="() => { $store.commit('setModalState', ModalState.NOTHING) }"
      />
      <div>
        <div> <span> {{ getDate() }} </span> </div>
        <div> <span> {{ $store.state?.openReview?.author }} </span> </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

h4 {
  font-size: 16px;
  margin: 0;
  font-weight: 600;
  margin-top: 0;
  margin-bottom: 2px;
  padding: 0;
  color: var(--text-color-2);
}

.modal-window {
  width: 60%;
  height: 80%;
  font-size: 16px;
  background: var(--panel-background-color);
  border-radius: 15px;
  padding: 25px;
}

.text {
  margin-top: 25px;
  height: 68%;
  overflow-y: scroll;
}

p {
  margin-top: 0;
  margin-right: 15px;
}

.top-block {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin-bottom: 15px
}

.bottom-block {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: end;
  margin-top: 3%;
}

.bottom-block span {
  color: var(--text-color-2);
}

</style>