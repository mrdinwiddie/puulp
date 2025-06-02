<script setup lang="ts">
import { ref } from 'vue'
import { NButton, NCard, NForm, NFormItem, NInput } from 'naive-ui'
const error = ref('')
const emit = defineEmits<{
  addClient: [newClient: string]
}>()
const newClient = ref('')
function formSubmit() {
  if (newClient.value.trim()) {
    emit('addClient', newClient.value)
    newClient.value = ''
  } else {
    error.value = 'invalid client name'
  }
}
</script>
<template>
  <NCard v-if="error">ERROR: {{ error }} </NCard>
  <NForm @submit.prevent="formSubmit" inline v-model="newClient">
    <NFormItem label="Add Client" path="count">
      <NInput
        type="text"
        placeholder="Client name..."
        v-model:value="newClient"
        @input="error = ''"
      />
    </NFormItem>
    <NFormItem>
      <NButton attr-type="button" @click="formSubmit" type="primary"> + </NButton>
    </NFormItem>
  </NForm>
</template>
