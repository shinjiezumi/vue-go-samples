<template>
  <v-app>
    <Header />
    <router-view />
    <Footer />
  </v-app>
</template>

<script>
// eslint-disable-next-line import/extensions,import/no-unresolved
import Header from './components/Header';
// eslint-disable-next-line import/extensions,import/no-unresolved
import Footer from './components/Footer';

export default {
  components: {
    Header,
    Footer,
  },
};
</script>
