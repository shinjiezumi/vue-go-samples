<template>
  <v-main>
    <v-row class="text-center" align="center" justify="center">
      <v-col class="mb-4" cols="12" sm="8">
        <h1>Searcher</h1>
      </v-col>
    </v-row>
    <v-row justify="center">
      <v-col cols="12" sm="6">
        <v-text-field @keydown.enter="search" v-model="keyword" outlined label="keyword"
                      append-icon="mdi-search-web"></v-text-field>
      </v-col>
    </v-row>
    <div v-if="this.isLoadingOn">
      <v-row justify="center">
        <v-progress-circular :size="50" color="primary" indeterminate/>
      </v-row>
    </div>
    <div v-else>
      <v-row v-if="this.error !== ''" class="text-center" justify="center">
        <v-alert type="error">{{ this.error }}</v-alert>
      </v-row>
      <v-row justify="center">
        <v-col class="mb-3" cols="12" xs="12" sm="8">
          <div v-for="result in this.searchResult" :key="result.id">
            {{result}}
          </div>
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
      keyword: ""
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
</style>