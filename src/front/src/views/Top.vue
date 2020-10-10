<template>
  <v-content>
    <v-row class="text-center mt-5 pa-3" align="center" justify="center">
      <IntroduceAppCard
          app-title="TodoList"
          app-sub-title="シンプルなTodoアプリです。"
          app-path="/todos"
      />
    </v-row>
  </v-content>
</template>

<script>
import { generateTitle } from "@/util";
import IntroduceAppCard from "../components/IntroduceAppCard.vue"

export default {
  name: "Top",
  title: generateTitle('トップページ'),
  components: {
    IntroduceAppCard
  }
}
</script>

<style scoped>

</style>