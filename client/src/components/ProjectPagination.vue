<script setup>
defineProps({
  count: Number,
  active: Number
})
</script>

<template>
  <div class="flex gap-x-2">
    <div
      v-for="index of count"
      :key="index"
      @click="$emit('onSelect', index)"
      class="py-1 px-3 bg-slate-200 rounded-lg cursor-pointer hover:bg-slate-300"
      :class="active === index && 'bg-slate-400'"
    >
      {{ index }}
    </div>
  </div>
</template>
