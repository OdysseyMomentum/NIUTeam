import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import { Auth0Plugin } from '@/auth/Auth0Plugin'
import { domain, clientId, cacheLocation, audience } from '../auth.config.json'
import Buefy from 'buefy'
import './assets/scss/app.scss'
import i18n from './i18n'
import '@mdi/font/css/materialdesignicons.css'
import vueFilterPrettyBytes from 'vue-filter-pretty-bytes'

Vue.use(vueFilterPrettyBytes)
Vue.use(Buefy)

Vue.use(
  Auth0Plugin, {
    domain: domain,
    clientId: clientId,
    audience: audience,
    cacheLocation: cacheLocation,
    onRedirectCallback: (appState: any) => {
      router.push(
        appState && appState.targetUrl ? appState.targetUrl : window.location.pathname
      )
    }
  }
)

Vue.config.productionTip = false

new Vue({
  router,
  store,
  i18n,
  render: h => h(App)
}).$mount('#app')


