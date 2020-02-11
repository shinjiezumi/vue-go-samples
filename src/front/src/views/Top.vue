<template>
  <v-content>
    <v-row class="text-center mt-5" align="center" justify="center">
      <v-col class="mb-4" cols="12" sm="8">
        <v-btn to="/todo" text>Todoリスト</v-btn>
      </v-col>
    </v-row>
    <v-row class="text-center" align="center" justify="center">
      <v-col>
        <v-btn to="/hello" text>Hello world</v-btn>
      </v-col>
    </v-row>
  </v-content>
</template>

<script>
  export default {
    name: "Top"
  }
</script>

<style scoped>

</style>