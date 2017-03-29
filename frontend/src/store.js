import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    siteInfo: {author: {}},
    currentUser: {}
  },
  mutations: {
    refresh: (state, info) => {
      state.info = info
    },
    signIn: (state, token) => {
      state.currentUser = {name: 'whoami'}
    },
    siteOut: state => {
      state.currentUser = {}
    }
  }
})

export default store
