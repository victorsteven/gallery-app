import Vue from 'vue'
import List from './components/Gallery/ListImages'
import VueMaterial from 'vue-material'
import 'vue-material/dist/vue-material.min.css'

Vue.use(VueMaterial)

Vue.config.productionTip = false

new Vue({
  render: h => h(List),
}).$mount('#app')
