<template>
  <v-main>
    <HelloWorld />
  </v-main>
</template>

<script>
// eslint-disable-next-line import/extensions,import/no-unresolved
import HelloWorld from '../components/HelloWorld';

export default {
  name: 'Hello',

  components: {
    HelloWorld,
  },

  data: () => ({
    //
  }),
};
</script>
