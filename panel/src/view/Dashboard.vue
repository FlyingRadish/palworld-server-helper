<template>
  <div class="flex-col-nowrap-center-start-center" style="width: 100%; max-width: 400px; margin-right: 8px;">
    <n-card title="状态">
      <template #header-extra>
        <div class="flex-row-nowrap-start-start-center">
          <div>{{ stateText }}</div>
          <div :style="stateStyle"></div>
        </div>
      </template>
      <div class="flex-col-nowrap-center-center-center">
        <n-progress type="dashboard" gap-position="bottom" :percentage="state.memory.usedPercent" />
        <div style="font-weight: 600;">内存</div>
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
import { computed, onMounted, onUnmounted, reactive } from 'vue'
import { MemStatus, PalPlayer } from '../model/PalApiModel'
import { getMemoryStatus, getOnlinePlayers, getServerState } from '../api/PalApi'


const state = reactive({
  timer: undefined as undefined | any,
  state: "unknown",
  onlinePlayers: {
    data: [] as Array<PalPlayer>,
  },
  memory: {
    status: {} as MemStatus,
    usedPercent: 0
  }
})

const stateText = computed(() => {
  switch (state.state) {
    case "running":
      return "运行中"
    case "rebooting":
      return "重启中"
    default:
      return "状态未知"
  }
})

const stateStyle = computed(() => {
  let color = "#999"
  switch (state.state) {
    case "running":
      color = "#00bc00"
      break
    case "rebooting":
      color = "#f1521c"
      break
    default:
      color = "#999"
      break
  }
  return {
    "background-color": color,
    "width": "8px",
    "height": "8px",
    "border-radius": "8px",
    "margin-left": "8px"
  }
})

const onRefreshState = async () => {
  try {
    let res = await getServerState()
    state.state = res.state
  } catch (error: any) {
  }
}

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
    onRefreshState()
    onRefreshMem()
    onRefreshOnlinePlayer()
  }, 5000)
}

onMounted(() => {
  startTimer()
})

onUnmounted(() => {
  if (state.timer) {
    clearInterval(state.timer)
  }
})

</script>
<style scoped></style>
