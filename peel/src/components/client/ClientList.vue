<script setup lang="ts">
import type { Client } from '@/types'

const props = defineProps<{
  clients: Client[]
}>()
</script>
<template>
  <UCard v-for="client in props.clients" :key="client.id">
    CLIENT
    <Client :client></Client>
  </UCard>
</template>
