<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { ref, h } from 'vue'
import {
  darkTheme,
  NConfigProvider,
  NCard,
  NSpace,
  NButton,
  NSwitch,
  NLayout,
  NLayoutSider,
} from 'naive-ui'
import type { MenuOption } from 'naive-ui'
import { NMenu } from 'naive-ui'

const heading = ref('PUULP | Metrics')
const theme = ref()
const collapsed = ref(false)
const menuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'home',
          },
        },
        { default: () => 'Home' },
      ),
    key: 'go-home',
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'about',
          },
        },
        { default: () => 'About' },
      ),
    key: 'go-about',
  },
]
</script>

<template>
  <NConfigProvider class="h-full" :theme="theme">
    <NSpace vertical>
      <header>
        <h1>{{ heading }}</h1>

        <!-- <NButton @click="theme = darkTheme"> Dark </NButton>
        <NButton @click="theme = null"> Light </NButton> -->
      </header>
      <!-- <NSwitch v-model:value="collapsed" /> -->
      <NLayout has-sider>
        <NLayoutSider
          bordered
          collapse-mode="width"
          :collapsed-width="64"
          :width="200"
          :collapsed="collapsed"
          show-trigger
          @collapse="collapsed = true"
          @expand="collapsed = false"
        >
          <n-menu :collapsed="collapsed" :collapsed-width="64" :options="menuOptions" />
        </NLayoutSider>
        <NLayout>
          <NSpace>
            <RouterView />
          </NSpace>
        </NLayout>
      </NLayout>
    </NSpace>
  </NConfigProvider>
</template>

<style scoped>
div {
  margin: 1rem auto;
}
nav {
  max-width: 100px;
}
span {
  margin: 2rem auto;
  padding-left: 2rem;
}
</style>
