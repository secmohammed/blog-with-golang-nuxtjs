import {
  mapGetters
} from 'vuex';
import Vue from 'vue'

const User = {
  install(Vue, options) {
    Vue.mixin({
      computed: {
        ...mapGetters('auth', {
          user: 'user',
          authenticated: 'authenticated',
        })
      }
    })
  }
}

Vue.use(User)