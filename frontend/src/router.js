import Vue from 'vue'
import Router from 'vue-router'

import auth from '@/engines/auth'
import site from '@/engines/site'

Vue.use(Router)

export default new Router({
  routes: [].concat(auth, site)
})
