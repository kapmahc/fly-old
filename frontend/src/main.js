require('./main.css')

import Vue from 'vue'
import Router from 'vue-router'
import I18n from 'vue-i18n'

Vue.use(I18n)
Vue.use(Router)

import App from './App'
import router from './router'
import {currentLocale, loadLocaleMessage} from './i18n'

Vue.config.productionTip = false

const locale = currentLocale()
const messages = {}
messages[locale] = {}
const i18n = new I18n({ locale, messages })

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  i18n,
  template: '<App/>',
  components: { App }
})

loadLocaleMessage(locale, (err, message) => {
  if (err) {
    console.error(err)
  } else {
    i18n.setLocaleMessage(locale, message)
  }
})
