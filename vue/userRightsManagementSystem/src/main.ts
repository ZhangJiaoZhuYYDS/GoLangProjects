import { createApp } from 'vue'
import { createPinia } from 'pinia'
import 'normalize.css'
import '@/assets/css/index.less'
import App from './App.vue'
import router from './router'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import { useLoginStore } from './stores/login/login'
import { localCache } from './utils/cache'
import { LOGIN_TOKEN } from './global/constants'

// import './assets/main.css'

const app = createApp(App)
app.use(createPinia())

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

const loginStore = useLoginStore()
loginStore.loadLocalCacheAction()

app.use(router)

app.mount('#app')
