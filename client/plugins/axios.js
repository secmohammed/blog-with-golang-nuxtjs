import { some, has} from 'lodash'
export default function({
    $axios,
    store,
    app
}) {
    $axios.onRequest(config => {
        config.url = `${process.env.BASE_URL}/${config.url}`
        if (store.getters['auth/authenticated']) {
            config.headers["Authorization"] = `Bearer ${store.getters['auth/token']}`                
        }

    })

    if (process.client) {
        $axios.onResponseError(error => {
            if (error.response.status == 422) {
                app.$toast.error(error.response.data.message || 'Invalid Data Supplied')
            }
            if (error.response.status == 500) {
                app.$toast.error('God ! , what did you do.. ')
            }
            if (error.response.status == 404) {
                app.$toast.error(error.response.data.message)
            }
            if (error.response.status == 401) {
                app.$toast.error(error.response.data.message)
            }
            return Promise.reject(error)
        })
        $axios.onRequest((response) => {
            let hasSuccess = some(response, res => has(res, "success"));
            let hasMessage = some(response, res => has(res, "message"));
            if (hasSuccess && hasMessage) {
                app.$toast.success(response.data.message);
            }
            if (!hasSuccess && hasMessage) {
                app.$toast.info(response.data.message);
            }
            return response;
        })
    }
}