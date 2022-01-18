<template>
  <v-form
      ref="form"
      lazy-validation
  >

    <v-text-field
        v-model="email"
        label="E-mail"
        required
    ></v-text-field>

    <v-text-field
        v-model="password"
        label="Password"
        required
    ></v-text-field>

    <v-checkbox
        label="Do you agree?"
        required
    ></v-checkbox>

    <v-btn
        class="mr-4"
        type="submit"
        @click="login"
    >
      submit
    </v-btn>
  </v-form>
</template>

<script>
import axios from "axios";

export default {
  name: "Login",
  data: function () {
    return {
      email: "",
      password: ""
    }
  },
  methods: {
    login() {
      axios({
        method: "POST",
        data: {
          email: this.email,
          password: this.password
        },
        url: "http://127.0.0.1:8090/login",
        headers: {
          // 'content-type': 'text/plain',
          'content-type': 'application/json'
        }})
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
