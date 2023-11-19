<template>
  <div id="app">
    <Header />
    <router-view />
  </div>
</template>

<script>
import Header from './components/PageHeader.vue';

export default {
  name: 'App',
  components: {
    Header
  }
};
</script>

<style>
/* Global styles */
</style>