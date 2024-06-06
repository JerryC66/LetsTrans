<template>
  <a-modal v-model:visible="visible"> </a-modal>
</template>

<script setup lang="ts">
  import { defineModel } from 'vue';

  const visible = defineModel<boolean>('visible', {
    type: Boolean,
    default: false,
    required: true,
  });
</script>

<style scoped></style>
