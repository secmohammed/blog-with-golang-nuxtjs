<template>
 <section class="container is-fluid">
      <div class="columns features">
        <div class="column is-4" v-for="post in posts" :key="post.ID">
          <div class="card is-shady">
            <div class="card-image" >
                <figure class="image is-4by3"> <!-- <figure class="image is-128x128">-->
                    <img src="https://ultraimg.com/images/yHrYA.jpg" alt="Placeholder image">
                </figure>
            </div>
            <div class="card-content">
              <div class="content">
                <h4> {{post.title}}</h4>
                <p>{{ post.body }}</p>
                <span class="button is-link"> more info....</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
    data() {
        return {
            posts: []
        }
    },
    async asyncData({ params, app, store }){
        let response = await app.$axios.$get('posts')
        store.commit("post/SET_POSTS", response)
        return {
            posts : response
        }
    },
}
</script>
