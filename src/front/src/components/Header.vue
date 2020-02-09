<template>
  <v-app-bar app color="primary" dark>
    <div class="d-flex align-center">
      <v-btn to="/" text>Vue Go Samples</v-btn>
    </div>
    <v-spacer />
    <v-btn to="/register" text>会員登録</v-btn>
    <v-btn to="/login" text>ログイン</v-btn>
  </v-app-bar>
</template>

<script>
  export default {
    name: 'Header',

    data: () => ({
      //
    }),
  };
</script>
