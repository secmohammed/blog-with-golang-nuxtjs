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
        <nuxt-link :to="{ name :'auth-login'}" class="navbar-item" v-if="!authenticated">
          Sign in
        </nuxt-link>
      </div>
    </div>
    </div>
  </nav>
</template>
<script>
  export default {
    
  }
</script>
