<script setup lang="ts">
import { ref } from 'vue'
import type { NavigationMenuItem } from '@nuxt/ui'
const items = ref<NavigationMenuItem[]>([
  {
    label: 'Home',
    // icon: 'i-lucide-circle-help',
    to: '/',
  },
  {
    label: 'About',
    // icon: 'i-lucide-circle-help',
    to: '/about',
  },
])
</script>
<template>
  <UNavigationMenu :items="items" class="w-full justify-center" />
</template>
