<script setup lang="ts">
import { defineProps, reactive } from 'vue';

const props = defineProps({
    info: RatedListItemInfo,
})

console.log(props.info)
</script>

<script lang="ts">
// move declarations to another file

export class RatedListItemInfo {
    rating: number;
    name: string;
    tags: string[];

    constructor (rating: number, name: string, tags: string[]) {
        this.rating = rating;
        this.name = name;
        this.tags = tags;
    }
}

function ratingToColor(rating: number): string {
    if (rating <= 2.5)
        return "#FF0000";
    else if (rating <= 5)
        return "#E3B064";
    else if (rating <= 7.5)
        return "#ABE364";
    else
        return "#85E364";
}
</script>

<template>
    <div class="rated-list-item">
        <div class="rated-list-item-rating-circle" :style="{ backgroundColor: ratingToColor(props.info?.rating!) }">
            <div class="rated-list-item-rating">
                {{ props.info?.rating.toFixed(1) }}
            </div>
        </div>

        <div class="rated-list-item-info-box">
            <div class="rated-list-item-name">
                {{ props.info?.name }}
            </div>
            <div class="rated-list-item-tags-box">
                <div class="rated-list-item-tag" v-for="(tag, idx) in props.info?.tags" :class="idx % 2 == 0 ? 'rated-list-item-tag-even' : 'rated-list-item-tag-odd'">
                    {{ tag }}
                </div>
            </div>
        </div>
    </div>
</template>

<style>
.rated-list-item {
    display: flex;
    align-items: center;
    margin-top: 10px;
    margin-bottom: 0px;
}

.rated-list-item-rating-circle {
    display: flex;
    align-items: center;
    width: 25px;
    aspect-ratio: 1 / 1;
    border-radius: 100%;
    padding: 5px;
    color: white;
    font-weight: 600;
    font-size: 15px;
    margin-right: 10px;
}

.rated-list-item-rating {
    width: 100%;
    text-align: center;
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
    align-items: center;
}

.rated-list-item-tag {
    display: flex;
    align-items: center;
    width: fit-content;
    padding-top: 3px;
    padding-bottom: 3px;
    padding-left: 10px;
    padding-right: 10px;
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