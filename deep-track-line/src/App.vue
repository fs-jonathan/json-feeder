<template>
  <div id="app">
    <h3>{{ errorMessage }}</h3>
    <!-- route outlet -->
    <!-- component matched by the route will render here -->
    <router-view></router-view>
  </div>
</template>

<script>
import Vue from 'vue'
import VueRouter from 'vue-router'
import axios from 'axios'

import Login from './components/Login.vue'
import GetRecord from './components/GetRecord.vue'

Vue.use(VueRouter)

const router = new VueRouter({
  routes: [
    { path: '/login', component: Login },
    { path: '/getRecord', component: GetRecord, props: { index: 0 } }
  ]
})

export default {
  router,
  name: 'app',
  data: () => ({
    errorMessage: ""
  }),
  mounted() {
    this.presetLiff();
    // this.initLiff();
    // this.initData("Hello");
    this.showData();
  },
  methods: {
    presetLiff: function() {
      // 設定済み
      if (document.getElementById('liff-line')) return;

      var liffScript = document.createElement('script');
      liffScript.src = 'https://d.line-scdn.net/liff/1.0/sdk.js';
      liffScript.id = 'liff-line';

      document.getElementByTagName('head')[0].appendChild(liffScript);
    },
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
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: left;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
