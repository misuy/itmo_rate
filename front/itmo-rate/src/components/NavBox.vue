<script setup lang="ts">
import NavItem from '@/components/NavItem.vue'
import HomeIcon from '@/components/icons/HomeIcon.vue'
import ClassIcon from '@/components/icons/ClassIcon.vue'
import AboutIcon from '@/components/icons/AboutIcon.vue'
import TeachersIcon from "@/components/icons/TeachersIcon.vue";
</script>

<template>
  <div class="nav-box">
    <NavItem :route-to="'/'">
      <template #icon>
        <HomeIcon />
      </template>
      <template #name>
        Главная
      </template>
    </NavItem>

    <NavItem :route-to="'/teachers'">
      <template #icon>
        <TeachersIcon />
      </template>
      <template #name>
        Преподаватели
      </template>
    </NavItem>

    <NavItem :route-to="'/subjects'">
      <template #icon>
        <ClassIcon />
      </template>
      <template #name>
        Предметы
      </template>
    </NavItem>

    <NavItem :route-to="'/about'">
      <template #icon>
        <AboutIcon />
      </template>
      <template #name>
        О сервисе
      </template>
    </NavItem>
  </div>
</template>

<style scoped>
.nav-box {
  display: flex;
  align-items: center;
  gap: 22px;
  flex-direction: column;
  width: 100%;
  margin-top: 15px;
}
</style>
