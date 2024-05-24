<script setup>
defineProps({
  id: String,
  title: String
})
</script>

<template>
  <div
    class="bg-white w-full p-2 border-b flex items-center gap-x-8 cursor-pointer hover:bg-slate-100"
  >
    <strong class="text-sm max-w-20 w-full">{{ id }}</strong>
    <h4 class="capitalize text-sm">{{ title }}</h4>
  </div>
</template>
