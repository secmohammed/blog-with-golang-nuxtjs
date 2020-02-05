<template>
  <nav class="navbar is-white">
    <div class="container">
    <div class="navbar-brand">
      <nuxt-link :to="{name : 'index' }" class="navbar-item">
        Blog
      </nuxt-link>
      <div class="navbar-burger burger" data-target="nav">
        <span></span>
        <span></span>
        <span></span>
      </div>
    </div>
    <div id="nav" class="navbar-menu">
      <div class="navbar-end">
        <template v-if="!authenticated">
          <nuxt-link :to="{ name :'auth-login'}" class="navbar-item">
            Sign in
          </nuxt-link>
          <nuxt-link :to="{ name :'auth-register'}" class="navbar-item">
            Sign up
          </nuxt-link>
        </template>
        <template v-else>
          <nuxt-link :to="{ name :'auth-change-password'}" class="navbar-item">
            Change Your Password
          </nuxt-link>
          
        </template>
      </div>
    </div>
    </div>
  </nav>
</template>
<script>
  export default {
    
  }
</script>
