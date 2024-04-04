<script setup lang="ts">
import { defineProps, type PropType } from 'vue';
import RatingCircle from './RatingCircle.vue';
import ChipComponent from "@/components/ChipComponent.vue";
import { useRouter } from 'vue-router';

interface Info {
  id: number;
  score: number;
  name: string;
  tags: string[];
}

const props = defineProps({
  pathPrefix: {
    type: String,
    default: "/"
  },
  info: Object as PropType<Info>,
});
const router = useRouter();
console.log(props.pathPrefix);

function openItemPage() {
  router.push(props.pathPrefix + props.info?.id);
}

</script>

<template>
  <div @click="openItemPage()" class="rated-list-item">
    <RatingCircle :rating="props.info?.score" :radius="18" />

    <div class="rated-list-item-info-box">
      <div class="rated-list-item-name">
        {{ props.info?.name }}
      </div>
      <div class="rated-list-item-tags-box">
        <ChipComponent
          style="font-size: 13px" 
          v-for="(tag, idx) in props.info?.tags" :key="tag" 
          :text="tag" 
          :color="idx % 2 == 0 ? 'yellow' : 'blue'"
        />
      </div>
    </div>
  </div>
</template>

<style>
.rated-list-item {
    display: flex;
    align-items: center;
    margin-bottom: 0px;
    gap: 15px;
    cursor: pointer;
    padding: 10px;
    border-radius: 15px;
    transition: .2s box-shadow;
}

.rated-list-item:hover {
  box-shadow: 0px 3px 10px 2px #f1f1f3;
  transition: .2s box-shadow;
 }

.rated-list-item-info-box {
    display: flex;
    flex-direction: column;
}

.rated-list-item-name {
    font-size: 18px;
    font-weight: 500;
    margin-bottom: 5px;
}

.rated-list-item-tags-box {
  display: flex;
  gap: 5px;
  align-items: center;
}

.rated-list-item-tag {
  display: flex;
  align-items: center;
  width: fit-content;
  padding: 3px 10px;
  margin-right: 7px;
  border: 1px solid black;
  font-weight: 500;
  font-size: 13px;
  border-radius: 12px;
}

.rated-list-item-tag-even {
    border-color: #FFE8AE;
    background-color: #FFF9E2;
    color: #7B6A2A;
}

.rated-list-item-tag-odd {
    border-color: #AEC5FF;
    background-color: #E2EAFF;
    color: #345B94;
}


</style>