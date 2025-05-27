<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { ref } from 'vue'
// import HelloWorld from './components/HelloWorld.vue'
import NavBar from './components/NavBar.vue'
const heading = ref('PUULP | Metrics')
</script>

<template>
  <main>
    <UApp>
      <header>
        <h1>{{ heading }}</h1>

        <div>
          <NavBar />
        </div>
      </header>
      <RouterView @createClient="createClient" />
    </UApp>
  </main>
</template>

<style scoped>
main {
  max-width: 1200px;
  margin: 1rem auto;
}
</style>
