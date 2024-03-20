<script setup lang="ts">
import { defineProps } from 'vue';

const props = defineProps({
    info: SubjectInfo,
})

console.log(props.info)
</script>

<script lang="ts">
// move declarations to another file
export class SubjectOwner {
    faculty: String;
    department: String;

    constructor (faculty: String, department: String) {
        this.faculty = faculty;
        this.department = department;
    }
}

export class SubjectInfo {
    rating: Number;
    name: String;
    owner: SubjectOwner;

    constructor (rating: Number, name: String, owner: SubjectOwner) {
        this.rating = rating;
        this.name = name;
        this.owner = owner;
    }
}
</script>

<template>
    <div class="subject">
        <div class="subject-rating">
            {{ props.info?.rating }}
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
    width: 30px;
    height: 30px;
    border: 1px solid black;
    border-radius: 100%;
    display: flex;
    align-items: center;
}

.subject-info-box {
    display: flex;
    flex-direction: column;
}

.subject-owner-box {
    display: flex;
    align-items: center;
}
</style>