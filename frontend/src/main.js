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

// init store data

import {get} from '@/ajax'
import {TOKEN} from '@/constants'

get('/site/info').then(function (rst) {
  document.title = rst.title
  store.commit('refresh', rst)
})
const token = sessionStorage.getItem(TOKEN)
if (token) {
  store.commit('signIn', token)
}
