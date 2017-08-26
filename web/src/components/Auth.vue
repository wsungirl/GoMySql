<template>
  <md-card>
    <md-card-header>
      <div class="md-title">Authenticate</div>
    </md-card-header>

    <md-card-content>
      <form @keyup.enter="sign">
        <md-input-container>
          <md-icon>account_circle</md-icon>
          <label>Name</label>
          <md-input v-model="name" type="text"></md-input>
        </md-input-container>

        <md-input-container>
          <md-icon>lock</md-icon>
          <label>Password</label>
          <md-input v-model="password" type="password"></md-input>
        </md-input-container>

        <md-button @click="sign" class="md-raised md-accent">Go deeper</md-button>
      </form>
    </md-card-content>
  </md-card>
</template>

<script>
var authObject = {
  user: {
    authenticated: false,
    token: '',
    id: 0
  },

  name: 'auth',
  data() {
    return {
      name: '',
      password: ''
    }
  },

  methods: {
    authenticate() {
      var vue = this

      this.$http.create({
        auth: {
          username: this.name,
          password: this.password
        }
      }).get(this.$dbapiEndpoint + '/users/' + authObject.user.id + '/auth')
        .then(function(response) {
          if (!response.data || !response.data.access_token) console.log(response)

          authObject.user.token = response.data.access_token
          window.localStorage.setItem('access_token', authObject.user.token)

          authObject.user.authenticated = true
          vue.$router.push({ name: 'dblist' })
        })
        .catch(function(error) {
          console.log(error)
        })
    },
    sign() {
      var vue = this

      this.$http.post(this.$dbapiEndpoint + '/users', { name: this.name, password: this.password })
        .then(function(response) {
          if (!response.data || !response.data.ID) console.log(response)

          authObject.user.id = response.data.ID
          vue.authenticate()
        })
        .catch(function(error) {
          console.log(error)
        })
    }
  }
}

export default authObject
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#app .md-card {
  margin: 8rem;
  padding: 2rem;
}
</style>
