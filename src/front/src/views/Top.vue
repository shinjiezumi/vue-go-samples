<template>
  <v-main>
    <v-row class="text-center mt-5 pa-3" align="center" justify="center">
      <v-col class="mb-4" cols="12" sm="4">
        <IntroduceAppCard
          app-title="TodoList"
          app-sub-title="シンプルなTodoアプリ"
          app-path="/todos"
          app-image="/img/TodoList.png"
        />
      </v-col>
      <v-col class="mb-4" cols="12" sm="4">
        <IntroduceAppCard
          app-title="Searcher"
          app-sub-title="複数サイトのコンテンツを検索するアプリ"
          app-path="/searcher"
          app-image="/img/Searcher.png"
        />
      </v-col>
    </v-row>
  </v-main>
</template>

<script>
// eslint-disable-next-line import/no-unresolved
import { generateTitle } from '@/util';
import IntroduceAppCard from '../components/IntroduceAppCard.vue';

export default {
  name: 'Top',
  title: generateTitle('トップページ'),
  components: {
    IntroduceAppCard,
  },
};
</script>

<style scoped>

</style>
