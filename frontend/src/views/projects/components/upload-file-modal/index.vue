<template>
  <a-modal v-model:visible="visible" width="auto">
    <a-upload action="/">
      <template #upload-button>
        <div
          style="
            background-color: var(--color-fill-2);
            color: var(--color-text-1);
            border: 1px dashed var(--color-fill-4);
            height: 158px;
            width: 380px;
            border-radius: 2;
            line-height: 158px;
            text-align: center;
          "
        >
          <div>
            Drag the file here or
            <span style="color: #3370ff"> Click to upload</span>
          </div>
        </div>
      </template>
    </a-upload>
  </a-modal>
</template>

<script setup lang="ts">
  import { ref, reactive, defineModel } from 'vue';

  const visible = defineModel<boolean>('visible', {
    type: Boolean,
    default: false,
    required: true,
  });
</script>

<style scoped></style>
