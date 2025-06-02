<script setup lang="ts">
import type { Client } from '@/types'
// import type { Client } from '../../types'
import { NCard } from 'naive-ui'
import ClientCard from '@/components/client/ClientCard.vue'
const props = defineProps<{
  clients: Client[]
}>()
</script>
<template>
  <div v-for="client in props.clients" :key="client.id">
    All Clients:
    <ClientCard :client></ClientCard>
  </div>
</template>
