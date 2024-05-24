<script setup>
import { ref } from 'vue'
import ProjectItem from './ProjectItem.vue'
import ProjectPagination from './ProjectPagination.vue'

// TODO: delete mock data
const projects = Array.from({ length: 50 }, (_, i) => ({
  id: `TGS-${i}`,
  title: `Project ${10 + i}`
}))

const data = ref(projects.slice(0, 20))
const activePage = ref(1)

const changePage = (newPage) => {
  const firstProject = (newPage - 1) * 20
  const lastProject = firstProject + 20

  data.value = projects.slice(firstProject, lastProject)
  activePage.value = newPage
}
</script>

<template>
  <div class="flex gap-x-14 flex-col border border-b-0">
    <div v-for="item in data" v-bind:key="item.id">
      <ProjectItem :id="item.id" :title="item.title" />
    </div>
  </div>

  <div class="flex justify-end pt-4">
    <ProjectPagination @on-select="changePage" :count="5" :active="activePage" />
  </div>
</template>
