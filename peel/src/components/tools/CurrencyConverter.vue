<script setup lang="ts">
import { NCard, NForm, NFormItem, NInputNumber, NSpace } from 'naive-ui'
import { ref, onBeforeMount, onUpdated } from 'vue'

const loading = ref(true)
const valueFro = ref(0)
const baseCurrencyRate = ref(0)
const toCurrencyRate = ref(0)
const data = ref(null)
const error = ref(null)
const url = ref('https://api.frankfurter.dev/v1/latest')
const options = ref<any>([])

const fetchData = async () => {
  data.value = null
  error.value = null
  try {
    await fetch(url.value)
      .then((res) => res.json())
      .then((json) => (data.value = json))
      .then(() => (loading.value = false))
      .catch((err) => (error.value = err))
  } catch (error) {
    console.error('Error fetching data:', error)
  }
}

onUpdated(() => {
  if (data.value && data.value['rates']) {
    baseCurrencyRate.value = 1
    toCurrencyRate.value = !!data.value ? data.value['rates']['USD'] : 1

    options.value = !!data.value ? data.value['rates'] : null
    console.log(options.value)
  }
})

onBeforeMount(() => {
  setTimeout(fetchData, 900)
})
</script>
<template>
  <NCard v-if="loading">loading... </NCard>
  <NCard variant="outline" v-else>
    Convert Currency
    <div>Rate: EUR {{ baseCurrencyRate }} -> USD {{ toCurrencyRate }}</div>
    <NForm inline v-model="valueFro">
      <NFormItem>
        <NSpace>
          <NInputNumber v-model:value="valueFro" name="valueFro" />
          ->
          {{ valueFro * toCurrencyRate }}
        </NSpace>
      </NFormItem>
    </NForm>
  </NCard>
</template>
