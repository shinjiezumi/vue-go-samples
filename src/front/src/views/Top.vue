<template>
  <v-content>
    <v-row class="text-center mt-5" align="center" justify="center">
      <v-col class="mb-4" cols="12" sm="8">
        <v-btn to="/todos" text class="text-none">TodoList</v-btn>
      </v-col>
    </v-row>
  </v-content>
</template>

<script>
  import { generateTitle } from "@/util";

  export default {
    title: generateTitle('トップページ'),
    name: "Top"
  }
</script>

<style scoped>

</style>