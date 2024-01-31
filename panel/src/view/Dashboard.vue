<template>
  <div class="flex-col-nowrap-center-start-center" style="width: 100%; max-width: 400px; margin-right: 8px;">
    <n-card title="内存状态">
      <div class="flex-row-nowrap-center-center-center">
        <n-progress type="dashboard" gap-position="bottom" :percentage="state.memory.usedPercent" />
      </div>
    </n-card>
    <n-card title="在线玩家">
      <div>
        <n-data-table :columns="palyerColumns" :data="state.onlinePlayers.data" style="height: 400px; width: 100%;" />
      </div>
    </n-card>
  </div>
</template>
<script setup lang="ts">
import { onMounted, onUnmounted, reactive } from 'vue'
import { MemStatus, PalPlayer } from '../model/PalApiModel'
import { getMemoryStatus, getOnlinePlayers } from '../api/PalApi'


const state = reactive({
  timer: undefined as undefined | any,
  onlinePlayers: {
    data: [] as Array<PalPlayer>,
  },
  memory: {
    status: {} as MemStatus,
    usedPercent: 0
  }
})



const onRefreshMem = async () => {
  try {
    state.memory.status = await getMemoryStatus()
    state.memory.usedPercent = Number((state.memory.status.used / state.memory.status.total * 100).toFixed(1))
  } catch (error: any) {
  }
}

const palyerColumns = [
  {
    title: "名称",
    key: "name"
  },
  {
    title: "uid",
    key: "uid"
  },
  {
    title: "steamId",
    key: "steamId"
  }
]

const onRefreshOnlinePlayer = async () => {
  try {
    state.onlinePlayers.data = await getOnlinePlayers()
  } catch (error: any) {
  }
}

const startTimer = () => {
  state.timer = setInterval(() => {
    onRefreshMem()
    onRefreshOnlinePlayer()
  }, 5000)
}

onMounted(()=> {
  startTimer()
})

onUnmounted(()=> {
  if (state.timer) {
    clearInterval(state.timer)
  }  
})

</script>
<style scoped></style>
