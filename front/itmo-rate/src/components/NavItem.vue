<template>
  <div class="nav-item">
    <div class="nav-item-icon">
      <slot name="icon"></slot>
    </div>

    <div class="nav-item-name">
      <slot name="name" class="nav-item-name"></slot>
    </div>
  </div>
</template>

<style>
.nav-item {
  display: flex;
  align-items: center;
  height: 25px;
  margin-top: 12px;
  margin-bottom: 12px;
  width: 100%;
}

.nav-item-icon {
  color: #656d7c;
  vertical-align: middle;
  height: 100%;
  width: auto;
}

.nav-icon {
    height: 100%;
    width: auto;
}

.nav-item-name {
  color: #656d7c;
  margin-left: 15px;
  vertical-align: middle;
  font-size: 15px;
}
</style>
