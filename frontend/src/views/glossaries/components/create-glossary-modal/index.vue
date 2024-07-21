<template>
  <a-modal
    v-model:visible="visible"
    :on-before-ok="postGlossaryTerm"
    :title="$t('glossary.create')"
    width="auto"
  >
    <div class="wrapper">
      <a-space direction="vertical" size="large">
        <a-input
          v-model="data.name"
          :style="{ width: '360px' }"
          placeholder="Please enter glossary name"
          allow-clear
        >
          <template #prefix>
            <icon-storage />
          </template>
        </a-input>

        <a-input
          v-model="data.author"
          :style="{ width: '360px' }"
          placeholder="Please enter author's name"
          allow-clear
        >
          <template #prefix>
            <icon-user />
          </template>
        </a-input>
        <a-input-group>
          <a-select
            v-model="data.source_lang"
            :style="{ width: '180px' }"
            placeholder="source language"
          >
            <template #prefix>
              <icon-subscribe-add />
            </template>
            <a-option>EN</a-option>
            <a-option>ZH</a-option>
            <a-option>JA</a-option>
          </a-select>
          <a-select
            v-model="data.target_lang"
            :style="{ width: '180px' }"
            placeholder="Target language"
          >
            <template #prefix>
              <icon-tag />
            </template>
            <a-option>EN</a-option>
            <a-option>ZH</a-option>
            <a-option>JA</a-option>
          </a-select>
        </a-input-group>
        <a-input
          v-model="data.comment"
          :style="{ width: '360px' }"
          placeholder="If any other comment..."
          allow-clear
        >
          <template #prefix>
            <icon-pen-fill />
          </template>
        </a-input>
      </a-space>
    </div>
  </a-modal>
</template>

<script setup lang="ts">
  import { defineModel, ref, reactive } from 'vue';
  import { useRoute, useRouter } from 'vue-router';
  import { Message } from '@arco-design/web-vue';
  import { useI18n } from 'vue-i18n';
  import { createGlossary } from '@/api/glossaries';

  const route = useRoute();
  const { t } = useI18n();
  const visible = defineModel<boolean>('visible', {
    type: Boolean,
    default: false,
    required: true,
  });

  const data = reactive({
    comment: '',
    author: '',
    name: '',
    source_lang: '',
    target_lang: '',
  });

  const postGlossaryTerm = async () => {
    try {
      console.log(data);
      const response = await createGlossary(data);
      if (response.data) {
        console.log('create glossary res:', response.data);
        Message.success(t('glossary.create.success'));
      }
    } catch (error) {
      Message.error(t('glossary.create.fail'));
      console.log(error);
    }
  };
</script>

<style scoped>
  .wrapper {
    display: flex;
    justify-content: center;
    margin: 20px;
  }
</style>
