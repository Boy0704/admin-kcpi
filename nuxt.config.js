export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  publicRuntimeConfig: {
    FileUrl: process.env.baseUrlFile,
    ApiKeyTiny: process.env.apiKeyTinyMCE 
  },

  router: {
    middleware: 'auth', // Terapkan middleware 'auth' pada semua rute
  },

  head: {
    title: 'admin-kcpi',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,400i,700&display=fallback' },
    ],
    script: [
      { src: '/plugins/jquery/jquery.min.js' },
      { src: '/plugins/bootstrap/js/bootstrap.bundle.min.js' },
      { src: '/disttemplate/js/adminlte.min.js' },
    ],
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
    '@/assets/plugins/fontawesome-free/css/all.min.css',
    '@/assets/disttemplate/css/adminlte.min.css',
    'quill/dist/quill.core.css',
    // for snow theme
    'quill/dist/quill.snow.css',
    // for bubble theme
    'quill/dist/quill.bubble.css'
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    { src: '~/plugins/vue-good-table', ssr: false },
    { src: '~plugins/nuxt-quill-plugin', ssr: false }
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    '@nuxtjs/auth',
        [
            'vue-toastification/nuxt',
            {
                timeout: 1000,
                draggable: false,
            },
        ],
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    // Workaround to avoid enforcing hard-coded localhost:3000: https://github.com/nuxt-community/axios-module/issues/308
    baseURL: process.env.baseUrlAPI,
  },

  auth: {
    strategies: {
      local: {
        endpoints: {
          login: {
            url: '/login',
            method: 'post',
            propertyName: 'data.token',
          },
          logout: false,
          user: {
            url: '/current-users',
            method: 'get',
            propertyName: 'data',
          },
        },
        // tokenRequired: true,
        // tokenType: 'bearer',
        // globalToken: true,
        // autoFetchUser: true,
      },
    },
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  }
}
