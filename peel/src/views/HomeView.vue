<script setup lang="ts">
import { ref } from 'vue'
import ClientForm from '@/components/client/ClientForm.vue'
import type { Client } from '@/types'
import ClientList from '@/components/client/ClientList.vue'
const clients = ref<Client[]>([])

function addClient(newClient: string) {
  clients.value.push({
    id: crypto.randomUUID(),
    name: newClient,
    sector: 'TEST',
  })
}
</script>
<!--
  mock login workflow
-->
<template>
  <main>
    <!-- creating a client should probably be in some sort of admin view -->
    <ClientForm @add-client="addClient"></ClientForm>
    <ClientList :clients></ClientList>
    <h4 v-if="!clients.length">Create a Client!</h4>
  </main>
</template>
