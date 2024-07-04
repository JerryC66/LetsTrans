<template>
  <a-modal v-model:visible="visible" width="auto">
    <a-upload @error="handleError" name="file" :custom-request="handleUpload">
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
  import type {
    FileItem,
    RequestOption,
    UploadRequest,
  } from '@arco-design/web-vue';

  const route = useRoute();
  const visible = ref(false);

  const { projectId } = route.params;
  const url = `https://api.lt.firlin.cn/projects/${projectId}/files`;

  const handleUpload = (option: RequestOption): UploadRequest => {
    const { onProgress, onError, onSuccess, fileItem } = option;
    const formData = new FormData();
    formData.append('files', fileItem.file);

    const xhr = new XMLHttpRequest();
    xhr.open('POST', url);

    xhr.upload.onprogress = function (event) {
      let percent;
      if (event.total > 0) {
        // 0 ~ 1
        percent = event.loaded / event.total;
      }
      onProgress(percent, event);
    };
    xhr.onload = () => {
      if (xhr.status < 200 || xhr.status >= 300) {
        onError(new Error(xhr.statusText));
      } else {
        onSuccess(JSON.parse(xhr.responseText));
      }
    };

    xhr.onerror = () => {
      onError(new Error('Upload failed'));
    };

    xhr.send(formData);

    return {
      abort() {
        xhr.abort();
      },
    };
  };

  const handleError = (fileItem: FileItem) => {
    console.error('Upload failed:', fileItem);
    alert(`File upload failed: ${fileItem.response}`);
  };
</script>

<style scoped></style>
