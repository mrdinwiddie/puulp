<script setup lang="ts">
import BrandCard from './BrandCard.vue'
import type { Brand } from '@/types'
import { NCard } from 'naive-ui'
const props = defineProps<{
  brands: Brand[]
}>()
</script>
<template>
  <NCard v-for="brand in props.brands" :key="brand.id">
    BRAND
    <BrandCard :brand></BrandCard>
  </NCard>
</template>
