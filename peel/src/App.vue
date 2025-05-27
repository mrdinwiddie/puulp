<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { ref } from 'vue'
import HelloWorld from './components/HelloWorld.vue'
import NavBar from './components/NavBar.vue'
const heading = ref('PUULP | Metrics')
</script>

<template>
  <main>
    <UApp>
      <!-- <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />

      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView /> -->
      <header>
        <h1>{{ heading }}</h1>

        <div>
          <NavBar />
          <!-- <nav>
            <RouterLink to="/">Home</RouterLink>
            <RouterLink to="/about">About</RouterLink>
          </nav> -->
        </div>
      </header>
      <RouterView />
    </UApp>
  </main>
</template>

<style scoped>
main {
  max-width: 1200px;
  margin: 1rem auto;
}
</style>
