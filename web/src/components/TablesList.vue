<template>
  <md-card>
    <md-card-header>
      <div class="md-title">DB Tables</div>
    </md-card-header>

    <md-card-content>

      <md-whiteframe>
        <md-input-container>
          <label>Table name</label>
          <md-input v-model="tablename"></md-input>
        </md-input-container>

        <md-input-container>
          <label>Columns (space separated)</label>
          <md-input v-model="column_names"></md-input>
        </md-input-container>

        <md-button :click="insert" class="md-raised md-accent">Insert</md-button>
      </md-whiteframe>

      <md-list>
        <md-list-item v-for="table in tables" :key="table.name">
          <router-link :to="'/db/'+$route.params.db_id+'/tables/'+table.name+'/rows'">{{table.name}}</router-link>
        </md-list-item>
      </md-list>

    </md-card-content>
  </md-card>
</template>

<script>
export default {
  name: 'tableslist',
  data() {
    return {
      tablename: '',
      column_names: '',
      tables: []
    }
  },
  methods: {
    created() {
      var vue = this

      this.$http.get(this.$dbapiEndpoint + '/dbs/' + this.$route.params.db_id + '/tables')
        .then(function(response) {
          if (!response.data) console.log(response)

          vue.tables = response.data
        })
        .catch(function(error) {
          console.log(error)
        })
    },
    insert() {
      var vue = this

      if (this.tanlename === '' || this.column_names === '') return

      var columnArr = this.column_names.split(' ')
      var columns = []

      for (var i = 0; i < columnArr.length; i++) {
        columns.push({ name: columnArr[i], type: 'string' })
      }

      this.$http.post(this.$dbapiEndpoint + '/dbs/' + this.$route.params.db_id + '/tables',
        { name: this.tablename, columns: this.columns })
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
