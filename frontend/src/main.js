require('./main.css')
// for bootstrap
window.$ = window.jQuery = require('jquery')
global.Tether = require('tether')
require('bootstrap')

import Vue from 'vue'

import App from '@/App'
import router from '@/router'
import store from '@/store'
import i18n from '@/i18n'

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  i18n,
  store,
  template: '<App/>',
  components: { App }
})
