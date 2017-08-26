<template>
  <div id="app">
    <md-whiteframe>
      <md-toolbar class="md-dense">
        <h2 class="md-title" style="flex: 1">
          <router-link to="/">DB Interface</router-link>
        </h2>

        <md-button @click="revoke" class="md-icon-button">
          <md-icon>exit_to_app</md-icon>
        </md-button>
      </md-toolbar>

    </md-whiteframe>

    <router-view></router-view>
  </div>
</template>

<script>
import axios from 'axios'
import Auth from '@/components/Auth'

export default {
  name: 'app',

  methods: {
    revoke() {
      var vue = this

      this.$http.get(this.$dbapiEndpoint + '/users/' + Auth.user.id + '/revoke')
        .then(function(response) {
          window.localStorage.removeItem('access_token')

          Auth.user.token = ''
          Auth.user.authenticated = false
          delete axios.defaults.headers.common['Authorization']

          vue.$router.push({ name: 'auth' })
        })
        .catch(function(error) {
          console.log(error)
        })
    }
  }
}
</script>

<style>
#app {
  font-family: 'Roboto', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

#app .md-toolbar {
  text-align: center;
}

#app .md-toolbar a {
  color: white;
  text-decoration: none;
}
</style>
