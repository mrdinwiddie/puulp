<script setup lang="ts">
import { ref } from 'vue'
import { NButton } from 'naive-ui'
const error = ref('')
const emit = defineEmits<{
  addClient: [newClient: string]
}>()
const newClient = ref('')
function formSubmit() {
  // error.value = ''
  if (newClient.value.trim()) {
    emit('addClient', newClient.value)
    newClient.value = ''
  } else {
    error.value = 'invalid client name'
  }
}
</script>
<template>
  <UCard v-if="error">ERROR: {{ error }} </UCard>
  <form @submit.prevent="formSubmit">
    <label>
      Add Client +
      <input v-model="newClient" name="newClient" @input="error = ''" />
    </label>
    <div>
      <button>ADD</button>
      <!-- <UButton>Add</UButton> -->
      <!-- <n-button>naive-ui</n-button> -->
    </div>
  </form>
</template>
