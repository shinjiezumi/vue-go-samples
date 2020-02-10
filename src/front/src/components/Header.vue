<template>
  <v-app-bar app color="primary" dark>
    <div class="d-flex align-center">
      <v-btn to="/" text>Vue Go Samples</v-btn>
    </div>
    <v-spacer/>
    <div v-if="!this.isLogin">
      <v-btn to="/register" text>会員登録</v-btn>
      <v-btn to="/login" text>ログイン</v-btn>
    </div>
    <div v-else>
      <v-btn text @click="logout">ログアウト</v-btn>
    </div>
  </v-app-bar>
</template>

<script>
  export default {
    name: 'Header',
    computed: {
      isLogin() {
        return this.$store.getters['auth/check']
      }
    },
    methods: {
      logout() {
        this.$store.commit("auth/setToken", "");
        this.$store.commit("auth/setUser", "");

        this.$router.push('/login');
      }
    }
  };
</script>
