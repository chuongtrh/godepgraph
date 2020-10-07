export default {
  // Target (https://go.nuxtjs.dev/config-target)
  // mode: 'spa',
  target: 'static',
  ssr: false,
  env: {
    baseUrl: process.env.BASE_URL || 'http://localhost:3200'
  },
  // Global page headers (https://go.nuxtjs.dev/config-head)
  head: {
    title: 'Godepgraph - A simple tool visualize Go package',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Global CSS (https://go.nuxtjs.dev/config-css)
  css: [
    'element-ui/lib/theme-chalk/index.css'
  ],

  // Plugins to run before rendering page (https://go.nuxtjs.dev/config-plugins)
  plugins: [
    '@/plugins/element-ui'
  ],

  // Auto import components (https://go.nuxtjs.dev/config-components)
  components: true,

  // Modules for dev and build (recommended) (https://go.nuxtjs.dev/config-modules)
  buildModules: [
  ],

  // Modules (https://go.nuxtjs.dev/config-modules)
  modules: [
    '@nuxtjs/axios',
  ],
  /*
    ** Axios module configuration
    ** See https://axios.nuxtjs.org/options
    */
  axios: {},
  // Build Configuration (https://go.nuxtjs.dev/config-build)
  build: {
    transpile: [/^element-ui/],
  }
}
