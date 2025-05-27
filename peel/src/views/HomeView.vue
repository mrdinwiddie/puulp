<script setup lang="ts">
import { ref } from 'vue'

const emit = defineEmits<{
  createClient: [newClient: string]
}>()

const newClient = ref('')

function createClient() {
  console.log(newClient.value)
  emit('createClient', newClient.value)
}
</script>

<template>
  <UCard>
    Home page
    <form @submit.prevent="createClient">
      <label>
        New Client
        <input v-model="newClient" name="newClient" />
      </label>
    </form>
  </UCard>
</template>
