<template>
  <v-col class="mb-4" cols="12" sm="4">
    <v-card class="mx-auto" elevation="5">
      <v-img src="/img/TodoList.png"/>
      <v-card-title>{{appTitle}}</v-card-title>
      <v-card-subtitle>{{appSubTitle}}</v-card-subtitle>
      <v-card-actions>
        <v-btn :to="appPath" color="orange lighten-2" text>View</v-btn>
      </v-card-actions>
    </v-card>
  </v-col>
</template>

<script>
export default {
  props: {
    "appTitle": {
      type: String,
    },
    "appSubTitle": {
      type: String,
    },
    "appPath": {
      type: String,
    },
  }
}


</script>