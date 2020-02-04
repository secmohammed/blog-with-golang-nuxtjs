
export default {
  mode: 'universal',
  /*
  ** Headers of the page
  */
  head: {
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  /*
  ** Customize the progress-bar color
  */
  loading: { color: '#fff' },
  /*
  ** Global CSS
  */
  css: [
  ],
  /*
  ** Plugins to load before mounting the App
  */
  toast: {
    position: 'top-right',
    duration: 800,
  },
  plugins: [
    './plugins/mixins/user.js',
    './plugins/axios.js',
    {
      src: './plugins/vee-validate.js',
      ssr: true
    }
  ],
  /*
  ** Nuxt.js dev-modules
  */
  buildModules: [
  ],
  /*
  ** Nuxt.js modules
  */
  modules: [
    // Doc: https://github.com/nuxt-community/modules/tree/master/packages/bulma
    '@nuxtjs/bulma',
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios',
    "@nuxtjs/toast",
    // Doc: https://github.com/nuxt-community/dotenv-module
    '@nuxtjs/dotenv',
  ],
  /*
  ** Axios module configuration
  ** See https://axios.nuxtjs.org/options
  */
   axios: {
    baseURL: 'http://localhost:8000/api/',
    proxy: true,
    redirectError: {
      401: '/auth/login',
      500: '/'
    },
  },
  auth: {
    redirect: {
      login: '/auth/login'
    },
    strategies: {
      local: {
        endpoints: {
          login: {
            url: '/auth/login',
            method: 'post',
            propertyName: false
          },
          user: {
            url: '/user',
            method: 'get',
            propertyName: 'data'
          },
          logout: {
            url: '/auth/logout',
            method: 'post',
            propertyName: 'data'
          },
        }
      }
    }
  },
  /*
  ** Build configuration
  */
  build: {
    postcss: {
      preset: {
        features: {
          customProperties: false
        }
      }
    },
    /*
    ** You can extend webpack config here
    */
    extend (config, ctx) {
    }
  }
}
