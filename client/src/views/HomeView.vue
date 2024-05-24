<script setup>
import ProjectList from '@/components/ProjectList.vue'

// TODO: delete mock data
const projects = Array.from({ length: 20 }, (_, i) => ({
  id: `TGS-${i}`,
  title: `Project ${10 + i}`
}))
</script>

<template>
  <h1 class="text-2xl mb-6 uppercase">Projects</h1>
  <ProjectList :data="projects" />
</template>
