<script setup lang="ts">
import { ref } from 'vue'
import BrandForm from '../brand/BrandForm.vue'
import type { Brand } from '@/types'
import BrandList from '../brand/BrandList.vue'
const brands = ref<Brand[]>([])
import type { Client } from '@/types'

function addBrand(newBrand: string) {
  console.log('t', newBrand)

  brands.value.push({
    id: crypto.randomUUID(),
    name: newBrand,
    sector: 'TEST2222',
  })
  console.log(brands.value)
}
const props = defineProps<{
  client: Client
}>()
</script>
<template>
  <!-- TODO
    add brands,
    brands are created under a client,
    metrics are added to brands,
    client metrics are metrics accross brands,
    plan data intake / ingestion / import
    plan data outtake / excretion / export
  -->
  <UCard>
    TEST2
    {{ props.client.name }}
  </UCard>
  <BrandForm @add-brand="addBrand"></BrandForm>
  <BrandList :brands></BrandList>
  <h4 v-if="!brands.length">Create a Brand!</h4>
</template>
