// js/getRecord.js
export default {
  name: "GetRecord",
  data () {
    return {
      loading: false,
      results: null,
      error: null
    }
  },
  created () {
    // fetch the data when the view is created and the data is already being observed
    this.getJson(0)
  },
  template: `
    <div class="post">
      <div v-if="results">
        <div v-for="result in results">
          <div class="column" v-on:click="getJson(result.id)">
            <div class="header">
              <div class="title">{{ result.title }}</div><div class="subtitle">{{ result.subtitle }}</div>
            </div>
            <div>
              <div class="cost">¥{{ result.cost }}</div>
            <div>
            <div class="header">
              <img src="/img/icon_arrow_right.png" class="arrow"/>
              <div class="subtext">¥{{ result.compare }}({{ result.rate }}%)</div>
            </div>
          </div>
        </div>
      </div>

      <div class="loading" v-if="loading">
        Loading...
      </div>

      <div v-if="error" class="error">
        {{ error }}
      </div>
    </div>
  `,
  watch: {
    // call again the method if the route changes
    '$route': 'getJson'
  },
  methods: {
    getJson(id) {
      this.loading = true

      axios.post('/getJson', { "message": id })
           .then(response => {
             this.results = response.data;
           })
           .catch(err => {
             this.error = err.toString();
           })
           .finally(() => this.loading = false)
    }
  }
};
