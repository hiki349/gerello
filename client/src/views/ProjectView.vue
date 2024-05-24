<script setup>
import { useRoute } from 'vue-router'

const { params } = useRoute()
</script>

<template>
  <h1>Project {{ params.id }}</h1>
</template>
