<script setup lang="ts">
import { ref } from 'vue'
import BrandForm from '../brand/BrandForm.vue'
import type { Brand } from '@/types'
import BrandList from '../brand/BrandList.vue'
const brands = ref<Brand[]>([])
import type { Client } from '@/types'
import { NCard } from 'naive-ui'

function addBrand(newBrand: string) {
  brands.value.push({
    id: crypto.randomUUID(),
    name: newBrand,
    sector: 'CPG',
  })
  console.log(brands.value)
}
const props = defineProps<{
  client: Client
}>()
</script>
<template>
  <NCard>
    Name:
    {{ props.client.name }}
  </NCard>
  <BrandForm @add-brand="addBrand"></BrandForm>
  <BrandList :brands></BrandList>
  <h4 v-if="!brands.length">Create a Brand!</h4>
</template>
