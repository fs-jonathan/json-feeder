window.onload = function (e) {
  buildVue();
};

function buildVue() {
  var result = new Vue({
    el: '#app',
    data: {
      results: []
    },
    mounted () {
      this.getJson(0);
    },
    methods: {
      getJson: function(id) {
        axios.post('/getJson', { "message": id })
             .then(response => {
               this.results = response.data;
             })
      }
    }
  });
}
