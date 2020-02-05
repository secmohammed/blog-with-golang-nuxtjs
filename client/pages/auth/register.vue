<template>
    <div class="section">
  <div class="container is-fluid">
    <div class="columns is-centered">
      <div class="column is-6">
        <h1 class="title is-4">Login</h1>
        <form action="#" @submit.prevent="login">
          <div class="field">
            <label class="label">Email</label>
            <div class="control">
              <input class="input" v-model="form.email" v-validate="'required|email'" name="email" type="email" placeholder="e.g. alex@codecourse.com">
              <p v-if="errors.has('email')" class="is-danger has-text-danger"> {{ errors.first('email') }}</p>
            </div>
          </div>
          <div class="field">
            <label class="label">Name</label>
            <div class="control">
              <input class="input" v-model="form.name" v-validate="'required|min:4'" name="name" type="text" placeholder="e.g. Mohammed Osama">
              <p v-if="errors.has('name')" class="is-danger has-text-danger"> {{ errors.first('name') }}</p>
            </div>
          </div>

          <div class="field">
            <label class="label">Password</label>
            <div class="control">
              <input class="input" v-model="form.password" type="password" name="password" v-validate="'required|min:8|max:32'">
              <p v-if="errors.has('password')" class="is-danger has-text-danger"> {{ errors.first('password') }}</p>

            </div>
          </div>

          <div class="field">
            <p class="control">
              <button class="button is-info is-medium" :disabled="isButtonDisabled">
                Login
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
          email : '',
          name: ''
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
        this.$store.dispatch('auth/register', this.form).then((user) => {
          this.$toast.success(`Alright, you can now login ${user.name}`)
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