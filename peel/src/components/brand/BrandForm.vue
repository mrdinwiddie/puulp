<script setup lang="ts">
import { ref } from 'vue'
const error = ref('')
const emit = defineEmits<{
  addBrand: [newBrand: string]
}>()
const newBrand = ref('')
function brandFormSubmit() {
  console.log(newBrand.value)
  if (newBrand.value.trim()) {
    emit('addBrand', newBrand.value)
    newBrand.value = ''
  } else {
    error.value = 'invalid brand name'
  }
}
</script>
<template>
  <UCard v-if="error">ERROR: {{ error }} </UCard>
  <form @submit.prevent="brandFormSubmit">
    <label>
      Add Brand +
      <input v-model="newBrand" name="newBrand" @input="error = ''" />
    </label>
    <div>
      <button>ADD</button>
    </div>
  </form>
</template>
