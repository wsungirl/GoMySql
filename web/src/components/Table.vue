<template>
  <md-card>
    <md-card-header>
      <div class="md-title">Table rows</div>
    </md-card-header>

    <md-card-content>

      <md-whiteframe>
        <md-input-container md-inline v-for="col in columns" :key="col">
          <label>{{col}}</label>
          <md-input :input="insertInput($index, $event)"></md-input>
        </md-input-container>

        <md-button :click="insert" class="md-raised md-accent">Insert</md-button>
      </md-whiteframe>

      <md-table>
        <md-table-header>
          <md-table-row>
            <md-table-head v-for="col in columns" :key="col">{{col}}</md-table-head>
          </md-table-row>
        </md-table-header>

        <md-table-body>
          <md-table-row v-for="row in rows" :key="row.id">
            <md-table-cell v-for="d in row.data" :key="d">{{d}}</md-table-cell>
          </md-table-row>
        </md-table-body>
      </md-table>

    </md-card-content>
  </md-card>
</template>

<script>
export default {
  name: 'table',
  data() {
    return {
      input: [],
      columns: ['name', 'data'],
      rows: [
        {
          id: 1,
          data: ['xyz', 2]
        },
        {
          id: 2,
          data: ['abc', 3]
        }
      ]
    }
  },
  methods: {
    created() {
      var vue = this

      this.$http.get(this.$dbapiEndpoint +
        '/dbs/' + this.$route.params.db_id +
        '/tables/' + this.$route.params.table_id +
        '/rows')
        .then(function(response) {
          if (!response.data) console.log(response)

          vue.rows = response.data.rows
          vue.columns = response.data.columns
        })
        .catch(function(error) {
          console.log(error)
        })
    },
    insert() {
      var vue = this

      if (this.input.length === 0) return

      this.$http.post(this.$dbapiEndpoint +
        '/dbs/' + this.$route.params.db_id +
        '/tables/' + this.$route.params.table_id +
        '/rows',
        this.input)
        .then(function(response) {
          if (!response.data || !response.data.ok) console.log(response)

          vue.created()
        })
        .catch(function(error) {
          console.log(error)
        })
    },
    insertInput(index, event) {
      this.input[index] = event.target.value
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
