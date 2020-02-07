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
    }
};
export const actions = {
    fetchPosts({ commit }) {
        return this.$axios.$get("posts").then(response => {
            commit("SET_POSTS", response);
            return response;
        });
    }
};
