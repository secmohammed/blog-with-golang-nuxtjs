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
    },
    token(state) {
      return state.token
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
  async changePassword({state}, payload) {
      const response = await this.$axios.post('auth/change-password', payload)
      console.log(response)
  },
  register(_, payload) {
    return new Promise((resolve, reject) => {
      this.$axios.$post('auth/register', {
        email: payload.email,
        password: payload.password,
        name: payload.name
      }).then((res) => resolve(res.user)).catch(err => reject(err))
    })
  },
  login({ commit }, payload) {
    return new Promise((resolve, reject) => {
      this.$axios.$post('auth/login', {
          email: payload.email,
          password: payload.password
     }).then((res) => {
         commit("SET_USER", res.user)
         resolve(res.user)
     }).catch(err => reject(err))

    })
  }
}