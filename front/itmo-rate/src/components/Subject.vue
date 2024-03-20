<script setup lang="ts">
import { l } from 'node_modules/vite/dist/node/types.d-FdqQ54oU';
import { defineProps } from 'vue';

const props = defineProps({
    info: SubjectInfo,
})

console.log(props.info)
</script>

<script lang="ts">
// move declarations to another file
export class SubjectOwner {
    faculty: string;
    department: string;

    constructor (faculty: string, department: string) {
        this.faculty = faculty;
        this.department = department;
    }
}

export class SubjectInfo {
    rating: number;
    name: string;
    owner: SubjectOwner;

    constructor (rating: number, name: string, owner: SubjectOwner) {
        this.rating = rating;
        this.name = name;
        this.owner = owner;
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
    <div class="subject">
        <div class="subject-rating" :style="{ backgroundColor: ratingToColor(props.info?.rating!) }">
            {{ props.info?.rating.toFixed(1) }}
        </div>

        <div class="subject-info-box">
            <div class="subject-name">
                {{ props.info?.name }}
            </div>
            <div class="subject-owner-box">
                <div class="subject-owner subject-faculty">
                    {{ props.info?.owner.faculty }}
                </div>
                <div class="subject-owner subject-department">
                    {{ props.info?.owner.department }}
                </div>
            </div>
        </div>
    </div>
</template>

<style>
.subject {
    display: flex;
    align-items: center;
}

.subject-rating {
    display: flex;
    align-items: center;
    width: fit-content;
    aspect-ratio: 1 / 1;
    border-radius: 100%;
    padding: 5px;
    color: white;
    font-weight: 600;
    font-size: 15px;
    margin-right: 10px;
}

.subject-info-box {
    display: flex;
    flex-direction: column;
}

.subject-name {
    font-size: 18px;
    font-weight: 500;
    margin-bottom: 5px;
}

.subject-owner-box {
    display: flex;
    align-items: center;
}

.subject-owner {
    display: flex;
    align-items: center;
    width: fit-content;
    padding-top: 4px;
    padding-bottom: 4px;
    padding-left: 10px;
    padding-right: 10px;
    margin-right: 7px;
    border: 1px solid black;
    font-weight: 500;
    font-size: 13px;
    border-radius: 12px;
}

.subject-faculty {
    border-color: #FFE8AE;
    background-color: #FFF9E2;
    color: #7B6A2A;
}

.subject-department {
    border-color: #AEC5FF;
    background-color: #E2EAFF;
    color: #345B94;
}


</style>