import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import store from './store'
import vuetify from './plugins/vuetify'


Vue.config.productionTip = false

new Vue({
  store,
  vuetify,
  VueRouter,
  render: h => h(App),
}).$mount('#app')
