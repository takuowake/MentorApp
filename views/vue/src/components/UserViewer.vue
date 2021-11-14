<template>
  <v-btn block
  @click="getData"
  >
    Block Button
  </v-btn>
</template>

<script>
import axios from 'axios'
export default {
  name: "UserViewer",
  data: function () {
    return {
      count: 0,
      username: ""
    }
  },
  methods: {
    getData() {
      axios({method: "GET", url: "http://127.0.0.1:8090/welcome",  headers: {"content-type": "text/plain" }})
          .then((data) => {
            // Signed in
            // const user = user;
            console.log(data);
            // ...
          })
          .catch((error) => {
            const errorCode = error.code;
            const errorMessage = error.message;
            console.log(errorCode);
            console.log(errorMessage);
          });
    }
  }
}
</script>

<style scoped>

</style>
