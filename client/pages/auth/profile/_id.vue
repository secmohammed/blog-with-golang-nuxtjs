<template>
    <img :src="avatar" v-if="hasAvatar" />
</template>
<script>
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
        hasAvatar() {
            return !!this.profile.avatar;
        },
        avatar() {
            return process.env.BASE_URL + this.profile.avatar;
        }
    }
};
</script>
