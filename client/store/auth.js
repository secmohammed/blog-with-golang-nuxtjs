export const state = () => ({
  loggedIn: false,
  user: null,
  token: null
});
export const getters = {
  authenticated(state) {
    return state.loggedIn;
  },

  user(state) {
    return state.user;
  },
  token(state) {
    return state.token;
  }
};
export const mutations = {
  SET_USER(state, payload) {
    state.user = payload;
    state.loggedIn = true;
    state.token = payload.token;
  },
  UNSET_USER(state) {
    state.user = null;
    state.loggedIn = false;
    state.token = null;
  }
};
export const actions = {
  updateProfile({ commit, state }, payload) {
    return new Promise((resolve, reject) => {
      this.$axios
        .$post(`auth/profile`, payload, {
          header: {
            "Content-Type": "multipart/form-data"
          }
        })
        .then(res => {
          commit("SET_USER", {
            ...res.user,
            token: state.user.token
          });
          resolve(res);
        })
        .catch(err => reject(err));
    });
  },
  resetPassword(_, payload) {
    return new Promise((resolve, reject) => {
      this.$axios
        .$post(`auth/reset-password/${payload.token}`, payload.form)
        .then(res => resolve(res.user))
        .catch(err => reject(err));
    });
  },
  forgetPassword(_, payload) {
    return new Promise((resolve, reject) => {
      this.$axios
        .$post("auth/forget-password", payload)
        .then(res => resolve(res.user))
        .catch(err => reject(err));
    });
  },
  changePassword(_, payload) {
    return new Promise((resolve, reject) => {
      this.$axios
        .$post("auth/change-password", payload)
        .then(res => resolve(res.user))
        .catch(err => reject(err));
    });
  },
  register(_, payload) {
    return new Promise((resolve, reject) => {
      this.$axios
        .$post("auth/register", {
          email: payload.email,
          password: payload.password,
          name: payload.name
        })
        .then(res => resolve(res.user))
        .catch(err => reject(err));
    });
  },
  logout({ commit }) {
    return new Promise((resolve, reject) => {
      this.$axios
        .$post("auth/logout")
        .then(res => {
          commit("UNSET_USER");
          resolve(res);
        })
        .catch(err => reject(err));
    });
  },
  login({ commit }, payload) {
    return new Promise((resolve, reject) => {
      this.$axios
        .$post("auth/login", {
          email: payload.email,
          password: payload.password
        })
        .then(res => {
          if (res.status) {
            commit("SET_USER", res.user);
            resolve(res.user);
          } else {
            resolve(res);
          }
        })
        .catch(err => reject(err));
    });
  }
};
