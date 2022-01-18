<template>
  <v-card
      color="grey lighten-4"
      flat
      height="60px"
      tile
  >
    <v-toolbar
        dense
        height="60px"
    >
      <v-toolbar-title>Title</v-toolbar-title>

      <v-btn
          v-for="link in links"
          :key="link"
          text
      >
        {{ link }}
      </v-btn>

      <v-icon>
        mdi-magnify
      </v-icon>
      <v-responsive max-width="260">
        <v-text-field
            dense
            flat
            hide-details
            rounded
            solo-inverted
        ></v-text-field>
      </v-responsive>

      <v-spacer></v-spacer>



      <v-btn>
        作成
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-email</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-bell</v-icon>
      </v-btn>

      <div v-if="username">
        <v-btn icon>
          <v-avatar color="indigo">
            <v-icon dark>
              mdi-account-circle
            </v-icon>
          </v-avatar>
        </v-btn>
      </div>
      <div v-else>
        <v-btn>ログインする</v-btn>
        <v-btn>登録する</v-btn>
      </div>
    </v-toolbar>
  </v-card>
</template>

<script>
export default {
  data: () => ({
    links: [
      '大学院生をさがす',
      '大学院生を募集する',
    ],
    username: ""
  }),
}
</script>
