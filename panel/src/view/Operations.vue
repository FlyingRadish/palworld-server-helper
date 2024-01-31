<template>
  <div class="flex-col-nowrap-center-start-center" style="width: 100%; max-width: 400px; margin-right: 8px;">
    <n-card title="广播">
      <template #header-extra>
        <n-space align="center">
          <div class="flex-row-nowrap-start-start-start">
            <div style="margin-right: 4px;">下划线替换空格</div>
            <n-switch v-model:value="state.broadcast.replaceBlank"></n-switch>
          </div>
          <n-button @click="onBoradcast">广播</n-button>
        </n-space>
      </template>
      <div>
        <n-input v-model:value="state.broadcast.content" type="textarea" style="width: 400px;" placeholder="游戏服务端暂不支持中文"/>
      </div>
    </n-card>

    <n-card title="RCON">
      <div>
        <n-input-group>
          <n-input v-model:value="state.rcon.command" style="width: 400px;" placeholder="请输入RCON命令" />
          <n-button @click="onRconSend">发送</n-button>
        </n-input-group>
        <n-input :value="state.rcon.response" type="textarea" style="margin-top: 4px;" placeholder="游戏服务端响应内容" />
      </div>
    </n-card>

    <n-card title="服务器管理">
      <div>
        <div class="flex-row-nowrap-start-start-center">
          <n-switch v-model:value="state.reboot.notifyReboot"></n-switch>
          <div style="flex-grow: 1; margin-left: 4px;">{{ rebootMode }}</div>
          <n-button @click="onReboot">重启</n-button>
        </div>
      </div>
    </n-card>
  </div>
</template>
<script setup lang="ts">

import {
  computed,
  reactive
} from "vue"
import { broadcast, rcon, reboot } from "../api/PalApi"
import { useMessage } from 'naive-ui'

const message = useMessage()

const state = reactive({
  broadcast: {
    content: "",
    replaceBlank: true
  },
  rcon: {
    command: "",
    response: ""
  },
  reboot: {
    notifyReboot: true
  }
})

const onBoradcast = async () => {
  try {
    let content = state.broadcast.content
    if (state.broadcast.replaceBlank) {
      content = content.replace(" ", "_")
    }
    await broadcast({content})
    message.info("广播成功")
  } catch (error: any) {
    message.error("网络错误")
  }
}

const onRconSend =  async () => {
  try {
    state.rcon.response = ""
    let res = await rcon({command: state.rcon.command})
    state.rcon.response = res.data
    message.info("RCON转发成功")
  } catch (error: any) {
    message.error("网络错误")
  }
}

const rebootMode = computed(() => {
  if (state.reboot.notifyReboot) {
    return "倒计时踢出在线玩家"
  } else {
    return "立即重启"
  }
})

const onReboot = async () => {
  try {
    await reboot({notifyReboot: state.reboot.notifyReboot})
    message.info("重启命令发送成功")
  } catch (error: any) {
    message.error("网络错误")
  }
}

</script>
<style scoped>
.item-card {}
</style>
