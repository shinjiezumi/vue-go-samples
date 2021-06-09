<template>
  <v-card class="mx-auto" elevation="5" outlined>
    <v-img :src="appImage" width="100%" height="100%" />
    <v-card-title>{{ appTitle }}</v-card-title>
    <v-card-subtitle class="appSubTitle">
      {{ appSubTitle }}
    </v-card-subtitle>
    <v-card-actions>
      <v-btn :to="appPath" color="orange lighten-2" text>
        View
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: 'IntroduceAppCard',
  props: {
    appTitle: {
      type: String,
    },
    appSubTitle: {
      type: String,
    },
    appPath: {
      type: String,
    },
    appImage: {
      type: String,
    },
  },
};
</script>

<style scoped>
.appSubTitle {
  height: 2rem;
  max-height: 2rem;
}
</style>
