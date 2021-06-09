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
        <v-text-field
          v-model="keyword"
          outlined
          label="keyword"
          append-icon="mdi-search-web"
          @keydown.enter="search"
        />
      </v-col>
    </v-row>
    <div v-if="isLoadingOn">
      <!--  ローディング  -->
      <v-row justify="center">
        <v-progress-circular :size="50" color="primary" indeterminate />
      </v-row>
    </div>
    <div v-else>
      <!--  エラーメッセージ  -->
      <v-row v-if="error !== ''" class="text-center" justify="center">
        <v-alert type="error" class="error-text">
          {{ error }}
        </v-alert>
      </v-row>
      <v-row justify="center">
        <v-col class="mb-3" cols="12" xs="12" sm="8">
          <!--  検索結果タブ  -->
          <v-tabs v-model="tabs" fixed-tabs>
            <v-tabs-slider />
            <v-tab :key="'Qiita'" :href="'#mobile-tabs-5-Qiita'" class="primary--text">
              <div class="site-title">
                Qiita
              </div>
            </v-tab>
            <v-tab :key="'SlideShare'" :href="'#mobile-tabs-5-SlideShare'" class="primary--text">
              <div class="site-title">
                SlideShare
              </div>
            </v-tab>
            <v-tab :key="'Feedly'" :href="'#mobile-tabs-5-Feedly'" class="primary--text">
              <div class="site-title">
                Feedly
              </div>
            </v-tab>
          </v-tabs>
          <!--  検索結果  -->
          <v-tabs-items v-model="tabs">
            <v-tab-item :key="'Qiita'" :value="'mobile-tabs-5-Qiita'">
              <QiitaList :items="searchResult.Qiita" />
            </v-tab-item>
            <v-tab-item :key="'SlideShare'" :value="'mobile-tabs-5-SlideShare'">
              <SlideList :items="searchResult.SlideShare" />
            </v-tab-item>
            <v-tab-item :key="'Feedly'" :value="'mobile-tabs-5-Feedly'">
              <FeedList :items="searchResult.Feedly" />
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

.error-text {
  white-space: pre-line;
  text-align: left;
}
</style>

<script>
// eslint-disable-next-line import/no-unresolved
import { generateTitle } from '@/util';
// eslint-disable-next-line import/extensions,import/no-unresolved
import QiitaList from '../components/searcher/QiitaList';
// eslint-disable-next-line import/extensions,import/no-unresolved
import SlideList from '../components/searcher/SlideList';
// eslint-disable-next-line import/extensions,import/no-unresolved
import FeedList from '../components/searcher/FeedList';

export default {
  name: 'Searcher',
  title: generateTitle('Searcher'),
  components: { QiitaList, SlideList, FeedList },
  data() {
    return {
      keyword: '',
      tabs: null,
    };
  },
  computed: {
    searchResult() {
      return this.$store.getters['searcher/getResult'];
    },
    isLoadingOn() {
      return this.$store.getters['loading/isOn'];
    },
    error() {
      return this.$store.getters['error/getError'];
    },
    errorCode() {
      return this.$store.getters['error/getCode'];
    },
  },
  created() {
    this.init();
  },
  methods: {
    init() {
      this.$store.dispatch('searcher/init');
    },
    search() {
      if (this.keyword === '' || this.isLoadingOn) return;

      const params = {
        q: this.keyword,
      };
      this.$store.dispatch('searcher/search', params);
    },
  },
};
</script>
