<script setup lang="ts">
import type { Brand } from '@/types'

const props = defineProps<{
  brands: Brand[]
}>()
</script>
<template>
  <UCard v-for="brand in props.brands" :key="brand.id">
    BRAND
    <Brand :brand></Brand>
  </UCard>
</template>
