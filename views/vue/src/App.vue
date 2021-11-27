<template>
  <v-app>
    <router-view>
      <Home/>
      <UserViewer/>
    </router-view>
  </v-app>
</template>

<script>
import Home from './components/Home.vue'
import UserViewer from './components/UserViewer.vue'

export default {
  name: 'App',
  components: {
    Home,
    UserViewer,
  }
}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
