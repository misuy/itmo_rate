<script setup lang="ts"></script>

<template>
    <div class="user-box">
        <div class="user-info-box">
            <img class="user-img" src="../assets/icons/UserIcon.svg" />
            <div class="user-name">
                <div class="user-name-surname">Ненов</div>
                <div class="user-name-name">Владислав</div>
            </div>
        </div>
        <button class="user-exit-btn">
            Выйти
        </button>
    </div>

</template>

<style>
.user-box {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    color: #656d7c;
}

.user-info-box {
    width: 100%;
    display: flex;
    align-items: center;
    height: 40px;
}

.user-img {
    height: 130%;
    width: auto;
}

.user-name {
    margin-left: 15px;
    font-size: 15px;
    display: flex;
    flex-direction: column;
    vertical-align: middle;
}

.user-exit-btn {
    margin-top: 15px;
    padding: 0;
    width: 100%;
    height: 35px;
    color: #656d7c;
    font-size: 13px;
    background-color: #EFF1F6;
    border: 1px solid #C3CDE1;
    border-radius: 10px;
}
</style>
