import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    username: null
  },
  getters: {
    username: state => state.username
  },
  mutations: {
    updateUsername(state, username) {
      state.username = username;
    }
  },
  actions: {

  },
  modules: {
  }
})
