export const state = () => ({
    posts: []
});
export const getters = {
    posts(state) {
        return state.posts;
    },
    postsCount(state) {
        return state.posts.length;
    }
};
export const mutations = {
    SET_POSTS(state, payload) {
        state.posts = payload;
    },
    UNSET_POST_FROM_POSTS(state, title) {
        state.posts = state.posts.filter(post => post.title !== title);
    }
};
export const actions = {
    fetchPosts({ commit }) {
        return this.$axios.$get("posts").then(response => {
            commit("SET_POSTS", response);
            return response;
        });
    },
    destroyPost({ commit }, title) {
        return this.$axios
            .$delete(`posts/${title}`)
            .then(response => {
                commit("UNSET_POST_FROM_POSTS", title);
                return response;
            })
            .catch(err => err);
    }
};
