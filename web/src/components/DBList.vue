<template>
  <md-card>
    <md-card-header>
      <div class="md-title">User databases</div>
    </md-card-header>

    <md-card-content>

      <md-whiteframe>
        <md-input-container>
          <label>Database name</label>
          <md-input v-model="db_name"></md-input>
        </md-input-container>

        <md-button @click="insert" class="md-raised md-accent">Insert</md-button>
      </md-whiteframe>

      <md-list>
        <md-list-item v-for="db in dbs" :key="db.id">
          <router-link :to="'/db/'+db.id+'/tables'">{{db.name}}</router-link>
        </md-list-item>
      </md-list>

    </md-card-content>
  </md-card>
</template>

<script>
export default {
  name: 'dblist',
  data() {
    return {
      db_name: '',
      dbs: []
    }
  },
  methods: {
    created() {
      var vue = this

      this.$http.get(this.$dbapiEndpoint + '/dbs')
        .then(function(response) {
          if (!response.data) console.log(response)

          vue.dbs = response.data
        })
        .catch(function(error) {
          console.log(error)
        })
    },
    insert() {
      var vue = this

      if (this.db_name === '') return

      this.$http.post(this.$dbapiEndpoint + '/dbs', { name: this.db_name })
        .then(function(response) {
          if (!response.data || !response.data.ok) console.log(response)

          vue.created()
        })
        .catch(function(error) {
          console.log(error)
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#app .md-card {
  margin: 2rem;
  padding: .25rem;
}


#app .md-card-content>.md-whiteframe {
  padding: 1rem 2rem;
  margin-bottom: 2rem;
}
</style>
