<template>
  <v-main>
    <v-row class="text-center" align="center" justify="center">
      <v-col class="mb-4" cols="12" sm="8">
        <h1>Searcher</h1>
      </v-col>
    </v-row>
    <!--  検索フィールド  -->
    <v-row justify="center">
      <v-col cols="12" sm="6">
        <v-text-field @keydown.enter="search" v-model="keyword" outlined label="keyword"
                      append-icon="mdi-search-web"></v-text-field>
      </v-col>
    </v-row>
    <div v-if="this.isLoadingOn">
      <!--  ローディング  -->
      <v-row justify="center">
        <v-progress-circular :size="50" color="primary" indeterminate/>
      </v-row>
    </div>
    <div v-else>
      <!--  エラーメッセージ  -->
      <v-row v-if="this.error !== ''" class="text-center" justify="center">
        <v-alert type="error">{{ this.error }}</v-alert>
      </v-row>
      <v-row justify="center">
        <v-col class="mb-3" cols="12" xs="12" sm="8">
          <!--  検索結果タブ  -->
          <v-tabs v-model="tabs" fixed-tabs>
            <v-tabs-slider></v-tabs-slider>
            <v-tab v-for="site in sites" :href="'#mobile-tabs-5-'+site" class="primary--text" :key="site">
              <div class="siteTitle">{{site}}</div>
            </v-tab>
          </v-tabs>
          <!--  検索結果  -->
          <v-tabs-items v-model="tabs">
            <v-tab-item v-for="site in sites" :key="site" :value="'mobile-tabs-5-' + site">
              <div v-for="data in searchResult[site]" :key="data.ID">
                {{ data.Title }}
              </div>
            </v-tab-item>
          </v-tabs-items>
        </v-col>
      </v-row>
    </div>
  </v-main>
</template>

<script>
import { generateTitle, STATUS_UNAUTHORIZED } from "@/util";

export default {
  name: "Searcher",
  title: generateTitle('Searcher'),
  data() {
    return {
      keyword: "",
      sites: ['Feedly', 'SlideShare', 'Qiita'],
      tabs: null,
    }
  },
  created() {
  },
  watch: {},
  computed: {
    searchResult() {
      return this.$store.getters['searcher/getResult']
    },
    isLoadingOn() {
      return this.$store.getters['loading/isOn']
    },
    error() {
      return this.$store.getters['error/getError']
    },
    errorCode() {
      return this.$store.getters['error/getCode']
    }
  },
  methods: {
    search() {
      const params = {
        q: this.keyword
      };

      (async () => {
        await this.$store.dispatch('searcher/search', params);
        if (this.error !== '') {
          return this.handleError();
        }
      })()
    },
    handleError() {
      if (this.errorCode === STATUS_UNAUTHORIZED) {
        this.$store.dispatch('auth/logout');
      }
    }
  },
}
</script>

<style scoped>
.siteTitle {
  text-transform: none;
}
</style>