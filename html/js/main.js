
const Login = () => import('./login.js')
const GetRecord = () => import('./getRecord.js')

// Router Instance
const router = new VueRouter({
  routes: [
    { path: '/login', component: Login },
    { path: '/getRecord', component: GetRecord, props: { index: 0 } }
  ]
})

// Root Instance
const result = new Vue({
  router,
  el: '#app',
  mounted () {
    this.initLiff();
    // 検証用コード
    // this.initData(null);
    // this.showLogin();
    // this.showData();
  },
  data: {
    errorMessage: ""
  },
  methods: {
    initLiff: function() {
      var vm = this; // keep reference of viewmodel object
      liff.init(function (data) {
        vm.initData(data);
      }, function() {
        vm.errorMessage = "Liff Error";
      });
    },
    initData: function(data) {
      if (data) {
        axios.post('/loginLiff', { "lineUserId": data.context.userId })
             .then(response => {
               this.showData();
             })
             .catch(err => {
               this.showLogin();
             })
      } else {
        this.errorMessage = "Wrong Page Access";
      }
    },
    showLogin: function() {
      router.replace('/login');
    },
    showData: function() {
      router.replace('/getRecord');
    }
  }
}).$mount(`#app`);
