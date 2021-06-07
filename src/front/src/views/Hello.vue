<template>
  <v-main>
    <HelloWorld/>
  </v-main>
</template>

<script>
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
