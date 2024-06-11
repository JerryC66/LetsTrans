<template>
  <a-modal v-model:visible="visible" width="auto">
    <a-upload :action="url" @error="handleError" name="file">
      <template #upload-button>
        <div
          style="
            background-color: var(--color-fill-2);
            color: var(--color-text-1);
            border: 1px dashed var(--color-fill-4);
            height: 158px;
            width: 380px;
            border-radius: 2px;
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
  import { ref } from 'vue';
  import { useRoute } from 'vue-router';
  import type { FileItem } from '@arco-design/web-vue';

  const route = useRoute();
  const visible = ref(false);

  const { projectId } = route.params;
  const url = `/projects/${projectId}/files`;

  const handleError = (fileItem: FileItem) => {
    console.error('Upload failed:', fileItem);
    alert(`File upload failed: ${fileItem.response}`);
  };
</script>

<style scoped></style>
