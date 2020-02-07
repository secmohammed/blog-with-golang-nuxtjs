<template>
    <div class="section">
  <div class="container is-fluid">
    <div class="columns is-centered">
      <div class="column is-6">
        <h1 class="title is-4">Update Profile</h1>
        <form action="#" @submit.prevent="updateProfile">
          <div class="field">
            <label class="label">Change Avatar</label>
            <div class="control">
              <input class="input" @change.prevent="handleFileUpload" name="avatar" type="file" >
            </div>
          </div>
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
            <p class="control">
              <button class="button is-info is-medium">
                Update Profile
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
          avatar : '',
          email: "",
          name: ""
        }
      }
    },
    methods : {
      handleFileUpload() {
        this.form.avatar = event.target.files[0];
      },
      updateProfile() {
        let formData = new FormData()
        formData.append("avatar", this.form.avatar)
        formData.append("email", this.form.email)
        formData.append("name", this.form.name)
        this.$store.dispatch("auth/updateProfile", formData).then((res) => {
          this.$toast.success(res.message)
          this.$route.replace({
            name: "index"
          })
        }).catch(err => console.log(err))
      }
    }
  }
</script>