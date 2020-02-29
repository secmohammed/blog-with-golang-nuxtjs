<template>
    <div class="conatiner">
        <section class="articles">
            <div
                class="column is-8 is-offset-2"
                v-for="post in posts"
                :key="post.id"
            >
                <div class="card article">
                    <div class="card-content">
                        <div class="media">
                            <div class="media-center">
                                <img
                                    v-if="authorHasAvatar(post)"
                                    :src="authorAvatar(post)"
                                    class="author-image"
                                    alt="Placeholder image"
                                />
                            </div>
                            <div class="media-content has-text-centered">
                                <p class="title article-title">
                                    {{ post.title }}
                                </p>
                                <p class="subtitle is-6 article-subtitle">
                                    <a href="#">{{ post.user.name }}</a> on
                                    {{ new Date(post.created_at) }}
                                </p>
                            </div>
                        </div>
                        <div class="content article-body">
                            <p>{{ post.body }}</p>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </div>
</template>
<script>
import { mapGetters } from "vuex";
export default {
    data() {
        return {
            posts: []
        };
    },
    computed: {
        authorHasAvatar() {
            return post => !!post.user.avatar;
        },
        authorAvatar() {
            return post => "http://localhost:8000/" + post.user.avatar;
        }
    },
    async asyncData({ params, app, store }) {
        let response = await app.$axios.$get("posts");
        store.commit("post/SET_POSTS", response);
        return {
            posts: response
        };
    }
};
</script>
<style>
html,
body {
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
        Oxygen-Sans, Ubuntu, Cantarell, "Helvetica Neue", sans-serif;
    font-size: 14px;
    background: #f0f2f4;
}
.navbar.is-white {
    background: #f0f2f4;
}
.navbar-brand .brand-text {
    font-size: 1.11rem;
    font-weight: bold;
}
.hero-body {
    background-image: url(https://upload.wikimedia.org/wikipedia/commons/thumb/f/f6/Plum_trees_Kitano_Tenmangu.jpg/1200px-Plum_trees_Kitano_Tenmangu.jpg);
    background-position: center;
    background-size: cover;
    background-repeat: no-repeat;
    height: 500px;
}
.articles {
    margin: 5rem 0;
    margin-top: -200px;
}
.articles .content p {
    line-height: 1.9;
    margin: 15px 0;
}
.author-image {
    position: absolute;
    top: -30px;
    left: 50%;
    width: 60px;
    height: 60px;
    margin-left: -30px;
    border: 3px solid #ccc;
    border-radius: 50%;
}
.media-center {
    display: block;
    margin-bottom: 1rem;
}
.media-content {
    margin-top: 3rem;
}
.article,
.promo-block {
    margin-top: 6rem;
}
div.column.is-8:first-child {
    padding-top: 0;
    margin-top: 0;
}
.article-title {
    font-size: 2rem;
    font-weight: lighter;
    line-height: 2;
}
.article-subtitle {
    color: #909aa0;
    margin-bottom: 3rem;
}
.article-body {
    line-height: 1.4;
    margin: 0 6rem;
}
.promo-block .container {
    margin: 1rem 5rem;
}
</style>
