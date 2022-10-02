import { createApp } from 'vue'
import App from './App.vue'
// import {ElMessage, ElMessageBox} from "element-plus";
import router from './router'

import "./element-variables.scss"

const app = createApp(App).use(router)
// installElementPlus(app)
app.mount('#app')
