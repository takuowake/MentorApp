import Vue from 'vue'
import VueRouter from 'vue-router'
import UserViewer from '../components/UserViewer.vue'
import Home from '../components/Home.vue'
import Login from "@/components/Login";

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/welcome',
        name: 'UserViewer',
        component: UserViewer
    },
    {
        path: '/login',
        name: 'Login',
        component: Login
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router
