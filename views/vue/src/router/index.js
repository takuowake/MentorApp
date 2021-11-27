import Vue from 'vue'
import VueRouter from 'vue-router'
import UserViewer from '../components/UserViewer.vue'
import Home from '../components/Home.vue'

Vue.use(VueRouter)

export default new VueRouter({
    routes: [
        {
            path: '/',
            component: Home
        },
        {
            path: '/welcome',
            component: UserViewer
        }
    ]
})
