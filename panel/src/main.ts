import { createApp } from "vue"
import "./style.css"
import "./assets/flex.css"
import "./assets/responsive.css"
import App from "./App.vue"
import {
  // create naive ui
  create,
  // component
  NButton,
  NCard,
  NInput,
  NInputGroup,
  NProgress,
  NSpace,
  NSwitch,
  NDataTable,
  NMessageProvider
} from "naive-ui"

const naive = create({
  components: [
    NButton,
    NInput,
    NInputGroup,
    NCard,
    NSwitch,
    NSpace,
    NProgress,
    NDataTable,
    NMessageProvider
],
})

const app = createApp(App)
app.use(naive)
app.mount("#app")
