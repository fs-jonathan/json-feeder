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
      axios.post('/getJson', { "message": 0 })
           .then(response => {
             // console.log(response);
             this.results = response.data;
           })
    }
  });
}
