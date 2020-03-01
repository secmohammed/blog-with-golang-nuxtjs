<template>
  <div class="columns">
    <div class="container profile">
      <div class="section profile-heading">
        <div class="columns is-mobile is-multiline">
          <div class="column is-2">
            <span class="header-icon user-profile-image">
              <img alt="" :src="authorAvatar" />
            </span>
          </div>
          <div class="column is-4-tablet is-10-mobile name">
            <p>
              <span class="title is-bold">{{ profile.user.name }}</span>
              <br />
            </p>
            <p class="tagline">
              The users profile bio would go here, of course. It could be two
              lines or more or whatever. We should probably limit the amount of
              characters to ~500 at most though.
            </p>
          </div>
          <div class="column is-2-tablet is-4-mobile has-text-centered">
            <p class="stat-val">{{ userPostsCount }}</p>
            <p class="stat-key">posts</p>
          </div>
          <div class="column is-2-tablet is-4-mobile has-text-centered">
            <p class="stat-val">3</p>
            <p class="stat-key">likes</p>
          </div>
        </div>
      </div>

      <div class="box" style="border-radius: 0px;">
        <!-- Main container -->
        <div class="columns">
          <div class="column is-10">
            <p class="control has-addons">
              <input
                class="input"
                placeholder="Search your liked properties"
                style="width: 100% !important"
                type="text"
              />
              <button class="button">
                Search
              </button>
            </p>
          </div>
        </div>
      </div>
      <div class="columns is-mobile">
        <div
          class="column is-3-tablet is-6-mobile"
          v-for="post in profile.posts"
          :key="post.id"
        >
          <div class="card">
            <div class="card-image">
              <figure class="image is-4by3">
                <img alt="" src="http://placehold.it/300x225" />
              </figure>
            </div>
            <div class="card-content">
              <div class="content">
                <span class="tag is-dark subtitle"
                  >#{{ post.id }} {{ post.title }}</span
                >
                <p>
                  {{ post.body }}
                </p>
              </div>
            </div>
            <footer class="card-footer">
              <a class="card-footer-item">Share</a>
              <a
                class="card-footer-item"
                v-if="authenticated && user.id === post.user_id"
                @click.prevent="removePost(post.title)"
                >Delete</a
              >
            </footer>
          </div>
          <br />
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { mapActions } from "vuex";
export default {
  data() {
    return {
      profile: null
    };
  },
  async asyncData({ params, app }) {
    let response = await app.$axios.$get(`/users/profile/${params.id}`);
    return {
      profile: response
    };
  },
  computed: {
    userPostsCount() {
      return this.profile.posts.length;
    },
    authorHasAvatar() {
      return this.profile.avatar;
    },
    authorAvatar() {
      return "http://localhost:8000/" + this.profile.user.avatar;
    }
  },
  methods: {
    ...mapActions("post", {
      destroyPost: "destroyPost"
    }),
    removePost(title) {
      this.destroyPost(title);
      this.profile.posts = this.profile.posts.filter(
        post => title != post.title
      );
    }
  }
};
</script>
