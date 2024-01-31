

<template>
  <n-message-provider>
  <div class="flex-col-nowrap-start-start-center" style="height: 100vh; width: 100vw; background-color: #f0f0f0;">
    <div class="header-container flex-row-nowrap-center-center-center">
      帕鲁世界服务器助手
    </div>
    <div class="flex-row-wrap-center-start-start" style="width: 100%; flex-grow: 1; ">
      <dashboard />
      <operations />
    </div>
  </div>
</n-message-provider>
</template>
<script setup lang="ts">
import Dashboard from "./view/Dashboard.vue"
import Operations from "./view/Operations.vue"
</script>

<style scoped>
.header-container {
  width: 100%;
  height: 48px;
  z-index: 10;
  color: #333;
  font-size: 24px;
  font-weight: 300;
  background-color: #fefefe;
  box-shadow: rgba(50, 50, 93, 0.25) 0px 6px 12px -2px, rgba(0, 0, 0, 0.3) 0px 3px 7px -3px;
}
</style>
