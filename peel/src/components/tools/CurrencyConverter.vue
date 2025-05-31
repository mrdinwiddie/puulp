<script setup lang="ts">
import { ref, onBeforeMount, onUpdated, onMounted } from 'vue'

const loading = ref(true)

const valueFro = ref(0)
const baseCurrencyRate = ref(0)
const toCurrencyRate = ref(0) //
// todo : cache this on page load
const data = ref(null)
const error = ref(null)
const url = ref('https://api.frankfurter.dev/v1/latest')
const items = ref(null)

function convertCurency() {
  // testing async
  data.value = null
  error.value = null
  fetch(url.value)
    .then((res) => res.json())
    .then((json) => (data.value = json))
    .then(() => (loading.value = false))
    .catch((err) => (error.value = err))
}

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

// testing lifecycle functionality
onUpdated(() => {
  // set rates
  if (data.value && data.value['rates']) {
    // rework logic so that this is set based on other rates
    // TODO get matrix of rates instead of just for EUR
    baseCurrencyRate.value = 1
    // might need to check for this key
    toCurrencyRate.value = !!data.value ? data.value['rates']['USD'] : 1

    items.value = !!data.value ? data.value['rates'] : null
    // Object.keys(data.value['rates'])
    console.log(items.value)
  }
})

onBeforeMount(() => {
  setTimeout(fetchData, 900)
  // setTimeout(convertCurency, 900)
})
onMounted(() => {
  // fetchData()
})
</script>
<!-- TODO
  display a graph/chart of conversion rates,
  record & display a history of converted values
-->
<template>
  <UCard v-if="loading">loading... </UCard>
  <UCard variant="outline" v-else>
    Convert Currency
    <div>Rates - EUR {{ baseCurrencyRate }} -> USD {{ toCurrencyRate }}</div>
    <!-- something with this select element or the filter seems to have slowed down the app -->
    <!-- USelect can't accept this object because the data is the key. will need something else. -->
    <USelect v-if="!loading" v-model="items" :items="items" class="w-48" />
    <UForm>
      <label>
        FRO :
        <UInput v-model="valueFro" name="valueFro" />
      </label>
      <label> TO : </label>
      {{ valueFro * toCurrencyRate }}
    </UForm>
  </UCard>
</template>
