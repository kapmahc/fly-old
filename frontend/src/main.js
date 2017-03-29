require('./main.css')

import Vue from 'vue'
import Router from 'vue-router'
import I18n from 'vue-i18n'

Vue.use(I18n)
Vue.use(Router)

import App from './App'
import routes from './routes'
import messages from './messages'

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router: new Router({routes}),
  i18n: new I18n(messages),
  template: '<App/>',
  components: { App }
})
