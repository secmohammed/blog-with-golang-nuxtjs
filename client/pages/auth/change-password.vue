<template>
    <div class="section">
  <div class="container is-fluid">
    <div class="columns is-centered">
      <div class="column is-6">
        <h1 class="title is-4">Change Password</h1>
        <form action="#" @submit.prevent="changePassword">
          <div class="field">
            <label class="label">Current Password</label>
            <div class="control">
              <input class="input" v-model="form.current_password" v-validate="'required|min:8|max:32'" name="current_password" type="password">
              <p v-if="errors.has('current_password')" class="is-danger has-text-danger"> {{ errors.first('current_password') }}</p>
            </div>
          </div>

          <div class="field">
            <label class="label">Password</label>
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
                Change Password
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
    middleware: 'auth',
    data(){
      return {
        form : {
          password : '',
          current_password : '',
          password_confirmation: ''
        }
      }
    },
    computed: {
      isButtonDisabled() {
        return !!this.errors.items.length
      },

    },
    methods : {
      changePassword() {
        this.$store.dispatch('auth/changePassword', this.form).then(() => {
          if (this.$route.query.redirect) {
            this.$router.replace(this.$route.query.redirect)
            return
          }
          this.$router.replace({
            name: 'index'
          })

        })
      }
    }
  }
</script>