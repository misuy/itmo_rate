<script lang="ts">
import { defineComponent } from 'vue'
import OutlineButton from "@/components/OutlineButton.vue"

export default defineComponent ({
  name: "UserBox",
  components: {OutlineButton},
  props: {
    authorized: Boolean,
    name: String,
  }
})

</script>

<template>
  <div v-if="authorized">
    <div class="user-box">
      <div class="user-info-box">
        <img class="user-img" src="../assets/icons/UserIcon.svg">
        <div class="user-name">
          <div class="user-name-surname">
            {{ name }}
          </div>
        </div>
      </div>
      <OutlineButton style="width: 78%;" text="Выйти" />
    </div>
  </div>
  <div v-else>
    <OutlineButton style="width: 78%; margin-bottom: 15px" text="Вход" />
    <OutlineButton style="width: 78%;" text="Регистрация" />
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
  margin-bottom: 20px;
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
  background-color: #eff1f6;
  border: 1px solid #c3cde1;
  border-radius: 10px;
}
</style>
