<template>
  <v-app>
    <Menu/>
    <router-view/>
    <Footers/>
  </v-app>
</template>

<script>
import Menu from './components/Menu.vue'
import Footers from './components/Footers.vue'

export default {
  name: 'App',
  components: {
    Menu,
    Footers
  }
}
</script>
