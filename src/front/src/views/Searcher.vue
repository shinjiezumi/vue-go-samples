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
            <v-tab :href="'#mobile-tabs-5-Qiita'" class="primary--text" :key="'Qiita'">
              <div class="site-title">Qiita</div>
            </v-tab>
            <v-tab :href="'#mobile-tabs-5-SlideShare'" class="primary--text" :key="'SlideShare'">
              <div class="site-title">SlideShare</div>
            </v-tab>
            <v-tab :href="'#mobile-tabs-5-Feedly'" class="primary--text" :key="'Feedly'">
              <div class="site-title">Feedly</div>
            </v-tab>
          </v-tabs>
          <!--  検索結果  -->
          <!--  タブ  -->
          <v-tabs-items v-model="tabs">
            <v-tab-item :key="'Qiita'" :value="'mobile-tabs-5-Qiita'">
              <FeedList :items="searchResult.Feedly"></FeedList>
            </v-tab-item>
            <v-tab-item :key="'SlideShare'" :value="'mobile-tabs-5-SlideShare'">
              <SlideList :items="searchResult.SlideShare"></SlideList>
            </v-tab-item>
            <v-tab-item :key="'Feedly'" :value="'mobile-tabs-5-Feedly'">
              <FeedList :items="searchResult.Feedly"></FeedList>
            </v-tab-item>
          </v-tabs-items>
        </v-col>
      </v-row>
    </div>
  </v-main>
</template>

<style scoped>
a {
  text-decoration: none
}

.site-title {
  text-transform: none;
}
</style>

<script>
import { generateTitle, STATUS_UNAUTHORIZED } from "@/util";
import SlideList from "../components/searcher/SlideList"
import FeedList from "../components/searcher/FeedList"

export default {
  name: "Searcher",
  title: generateTitle('Searcher'),
  components: {SlideList, FeedList},
  data() {
    return {
      keyword: "",
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
      if (this.isLoadingOn)
        return

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
