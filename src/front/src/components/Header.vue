<template>
  <div>
    <v-app-bar app color="primary" dark>
      <v-app-bar-nav-icon @click="drawer = true" />
      <div class="d-flex align-center">
        <v-btn to="/" text>
          Vue Go Samples
        </v-btn>
      </div>
      <v-spacer />
    </v-app-bar>
    <v-navigation-drawer v-model="drawer" absolute temporary>
      <v-list v-if="!this.isLogin" dense>
        <v-list-item>
          <v-list-item-content>
            <v-btn to="/register" text>
              会員登録
            </v-btn>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-btn to="/login" text>
              ログイン
            </v-btn>
          </v-list-item-content>
        </v-list-item>
      </v-list>
      <v-list v-else dense>
        <v-list-item>
          <v-list-item-content>
            <v-btn text @click="logout">
              ログアウト
            </v-btn>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>

<script>
export default {
  name: 'Header',
  data() {
    return {
      drawer: false,
    };
  },
  computed: {
    isLogin() {
      return this.$store.getters['auth/check'];
    },
  },
  methods: {
    logout() {
      this.$store.dispatch('auth/logout');
    },
  },
};
</script>
