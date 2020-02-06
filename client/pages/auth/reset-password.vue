<template>
    <div class="section">
  <div class="container is-fluid">
    <div class="columns is-centered">
      <div class="column is-6">
        <h1 class="title is-4">Reset Password</h1>
        <form action="#" @submit.prevent="resetPassword">
          <div class="field">
            <label class="label">New Password</label>
            <div class="control">
              <input class="input" v-model="form.password" type="password" ref="password" name="password" v-validate="'required|min:8|max:32'">
              <p v-if="errors.has('password')" class="is-danger has-text-danger"> {{ errors.first('password') }}</p>

            </div>
          </div>
          <div class="field">
            <label class="label">Password Confirmation</label>
            <div class="control">
              <input class="input" v-model="form.password_confirmation" type="password" name="password_confirmation" v-validate="'required|min:8|max:32|confirmed:password'">
              <p v-if="errors.has('password_confirmation')" class="is-danger has-text-danger"> {{ errors.first('password_confirmation') }}</p>

            </div>
          </div>

          <div class="field">
            <p class="control">
              <button class="button is-info is-medium" :disabled="isButtonDisabled">
                Reset Password
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
          password : '',
          password_confirmation: ''
        },
        token: ''
      }
    },
    created() {
      this.token = this.$router.params.token
    },
    computed: {
      isButtonDisabled() {
        return !!this.errors.items.length
      },

    },
    methods : {
      resetPassword() {
        this.$store.dispatch('auth/resetPassword', {
          form:  this.form,
          token: this.token
        }).then(() => {
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