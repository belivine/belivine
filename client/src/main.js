import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import Vuex from 'vuex'
import { createProvider } from './vue-apollo'

Vue.use(Vuex)

Vue.use(require('vue-cookies'))
// set default config
Vue.$cookies.config('7d')

Vue.config.productionTip = false

new Vue({
  router,
  vuetify,
  apolloProvider: createProvider(),
  render: h => h(App)
}).$mount('#app')
