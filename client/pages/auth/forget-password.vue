<template>
    <div class="section">
  <div class="container is-fluid">
    <div class="columns is-centered">
      <div class="column is-6">
        <h1 class="title is-4">Forget Password</h1>
        <form action="#" @submit.prevent="login">
          <div class="field">
            <label class="label">Email</label>
            <div class="control">
              <input class="input" v-model="form.email" v-validate="'required|email'" name="email" type="email" placeholder="e.g. alex@codecourse.com">
              <p v-if="errors.has('email')" class="is-danger has-text-danger"> {{ errors.first('email') }}</p>
            </div>
          </div>
          <div class="field">
            <p class="control">
              <button class="button is-info is-medium" :disabled="isButtonDisabled">
                Send Me An Email With Token
              </button>
            </p>
          </div>
        </form>
      </div>
    </div>
  </div>
</div>
</template>
<script>
  import { mapActions, mapGetters} from 'vuex'
  export default {
    middleware: 'guest',
    data(){
      return {
        form : {
          email : ''
        }
      }
    },
    computed: {
      isButtonDisabled() {
        return !!this.errors.items.length
      },

    },
    methods : {
      login() {
        this.$store.dispatch('auth/forgetPassword', this.form).then(() => {
          if (this.$route.query.redirect) {
            this.$router.replace(this.$route.query.redirect)
            return
          }
          this.$router.replace({
            name: 'auth-login'
          })

        })
      }
    }
  }
</script>