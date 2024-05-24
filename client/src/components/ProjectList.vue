<script setup>
import ProjectItem from './ProjectItem.vue';

// TODO: create pagination
defineProps({
  data: [
    {
      id: String,
      title: String
    }
  ]
})
</script>

<template>
  <div class="flex gap-x-14 flex-col border border-b-0">
    <div v-for="item in data" v-bind:key="item.id">
      <ProjectItem :id="item.id" :title="item.title" />
    </div>
  </div>
</template>
