<script setup>
import { RouterView } from 'vue-router'
import ProjectList from './components/ProjectList.vue'

// TODO: delete mock data
const projects = Array.from({ length: 20 }, (_, i) => ({
  id: `TGS-${i}`,
  title: `Project ${10 + i}`
}))
</script>

<template>
  <div class="max-w-7xl m-auto p-8">
    <ProjectList :data="projects" />

    <RouterView />
  </div>
</template>
