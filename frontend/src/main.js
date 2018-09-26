import Vue from 'vue'
import App from './App'
import router from './router'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import autofocus from 'vue-autofocus-directive'
var $ = require('jquery')

window.jQuery = $

Vue.directive('autofocus', autofocus)
Vue.use(BootstrapVue)

Vue.config.productionTip = false
// Vue.config.devtools = true
/* eslint-disable no-new */

new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
