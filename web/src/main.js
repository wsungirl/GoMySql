// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import 'roboto-fontface/css/roboto/roboto-fontface.css'
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import 'vue-material/dist/vue-material.css'

import Vue from 'vue'
import VueMaterial from 'vue-material'
import axios from 'axios'

import App from './App'
import router from './router'
import Auth from '@/components/Auth'

Vue.config.productionTip = false

Vue.use(VueMaterial)

Vue.material.registerTheme('default', {
  primary: 'blue',
  accent: {
    color: 'orange',
    textColor: 'white'
  }
})

axios.defaults.headers.common['Content-Type'] = 'application/json'

Vue.prototype.$http = axios
Vue.prototype.$dbapiEndpoint = 'http://localhost:6060'

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App },
  beforeCreate() {
    Auth.user.accessToken = localStorage.getItem('access_token')
    Auth.user.authenticated = !!Auth.user.accessToken
    if (Auth.user.authenticated) axios.defaults.headers.common['Authorization'] = 'Bearer ' + Auth.user.accessToken
  }
})
