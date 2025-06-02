<script setup lang="ts">
import { NButton, NCard, NForm, NFormItem, NInput } from 'naive-ui'
import { ref } from 'vue'
const error = ref('')
const emit = defineEmits<{
  addBrand: [newBrand: string]
}>()
const newBrand = ref('')
function brandFormSubmit() {
  console.log('TEST', newBrand.value)
  if (newBrand.value.trim()) {
    emit('addBrand', newBrand.value)
    newBrand.value = ''
  } else {
    error.value = 'invalid brand name'
  }
}
</script>
<template>
  <NCard v-if="error">ERROR: {{ error }} </NCard>
  <NForm @submit.prevent="brandFormSubmit" inline v-model="newBrand">
    <NFormItem>
      Add Brand +
      <NInput
        type="text"
        placeholder="Brand name..."
        v-model:value="newBrand"
        name="newBrand"
        @input="error = ''"
      />
    </NFormItem>
    <NFormItem>
      <NButton attr-type="button" @click="brandFormSubmit" type="primary">ADD</NButton>
    </NFormItem>
  </NForm>
</template>
