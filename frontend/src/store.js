import Vue from 'vue'
import Vuex from 'vuex'
import jwtDecode from 'jwt-decode'

Vue.use(Vuex)

import {TOKEN} from '@/constants'

const store = new Vuex.Store({
  state: {
    siteInfo: {author: {}, languages: []},
    currentUser: {}
  },
  mutations: {
    refresh: (state, info) => {
      state.siteInfo = info
    },
    signIn: (state, token) => {
      try {
        state.currentUser = jwtDecode(token)
        sessionStorage.setItem(TOKEN, token)
      } catch (e) {
        console.error(e)
        sessionStorage.clear()
        state.currentUser = {}
      }
    },
    signOut: state => {
      sessionStorage.clear()
      state.currentUser = {}
    }
  }
})

export default store
