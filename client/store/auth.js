export const state = () => ({
    loggedIn: false,
    user: null,
    token: null
})
export const getters = {
    authenticated(state) {
      return state.loggedIn
    },

    user(state) {
      return state.user
    }
}
export const mutations = {
    SET_USER(state, payload) {
        state.user = payload
        state.loggedIn = true
        state.token = payload.token
    }
}
export const actions = {
  register(_, payload) {
    return new Promise((resolve, reject) => {
      this.$axios.$post('http://localhost:8000/api/auth/register', {
        email: payload.email,
        password: payload.password,
        name: payload.name
      }).then((res) => resolve(res.user)).catch(err => reject(err))
    })
  },
  login({ commit }, payload) {
    return new Promise((resolve, reject) => {
      this.$axios.$post('http://localhost:8000/api/auth/login', {
          email: payload.email,
          password: payload.password
     }).then((res) => {
         commit("SET_USER", res.user)
         resolve(res.user)
     }).catch(err => reject(err))

    })
  }
}