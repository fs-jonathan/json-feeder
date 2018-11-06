import Vue from 'vue'
import App from './App.vue'

// BootstrapVueの読み込み
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import 'tailwindcss/dist/tailwind.min.css'

Vue.config.productionTip = false

// BootstrapVueの使用
Vue.use(BootstrapVue)

new Vue({
  render: h => h(App),
}).$mount('#app')
