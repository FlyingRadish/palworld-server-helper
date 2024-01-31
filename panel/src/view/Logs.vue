<template>
  <div class="flex-col-nowrap-center-start-center" style="width: 100%; max-width: 600px; margin-right: 8px;">
    <n-card title="助手日志">
      <template #header-extra>
        <n-button @click="onRefreshLog">刷新</n-button>
      </template>
      <div class="flex-row-nowrap-center-center-center">
        <n-input :value="state.log.content" type="textarea" style="width: 100%; min-height:600px" placeholder="" />
      </div>
    </n-card>
  </div>
</template>
<script setup lang="ts">
import { reactive } from 'vue'



const state = reactive({
  log: {
    content: ""
  }
})



const onRefreshLog = () => {
  
}

</script>
<style scoped></style>
