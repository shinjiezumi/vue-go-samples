<template>
  <v-app>
    <Header />
    <router-view />
    <Footer />
  </v-app>
</template>

<script>
  import Header from "./components/Header";
  import Footer from "./components/Footer";
  export default {
    components: {
      Header: Header,
      Footer: Footer
    }
  }
</script>
