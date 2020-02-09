<template>
  <v-content>
    <HelloWorld/>
  </v-content>
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
