<script setup lang="ts">
import type { Client } from '@/types'

const props = defineProps<{
  clients: Client[]
}>()
</script>
<template>
  <UCard v-for="client in props.clients" :key="client.id">
    {{ client.name }}
  </UCard>
</template>
