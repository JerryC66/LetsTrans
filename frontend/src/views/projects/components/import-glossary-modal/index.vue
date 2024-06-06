<template>
  <a-modal
    v-model:visible="visible"
    :on-before-ok="postGlossaryTerm"
    :title="$t('translate.addterm')"
    width="auto"
  >
    <div class="wrapper">
      <a-space direction="vertical" size="large">
        <a-input-group>
          <a-input
            v-model="data.source_lang"
            :style="{ width: '180px' }"
            placeholder="source language"
          >
            <template #prefix>
              <icon-subscribe-add />
            </template>
          </a-input>
          <a-input
            v-model="data.source_text"
            :style="{ width: '280px' }"
            placeholder="Please enter source  text"
            allow-clear
          >
          </a-input>
        </a-input-group>
        <a-input-group>
          <a-input
            v-model="data.target_lang"
            :style="{ width: '180px' }"
            placeholder="Target language"
          >
            <template #prefix>
              <icon-tag />
            </template>
          </a-input>
          <a-input
            v-model="data.target_text"
            :style="{ width: '280px' }"
            placeholder="Please enter target text"
            allow-clear
          >
          </a-input>
        </a-input-group>
      </a-space>
    </div>
  </a-modal>
</template>

<script setup lang="ts">
  import { defineModel, ref, reactive } from 'vue';
  import { addTermToGlossary } from '@/api/glossaries';
  import { useRoute, useRouter } from 'vue-router';
  import { Message } from '@arco-design/web-vue';
  import { useI18n } from 'vue-i18n';

  const route = useRoute();
  const { t } = useI18n();
  const visible = defineModel<boolean>('visible', {
    type: Boolean,
    default: false,
    required: true,
  });

  const data = reactive({
    source_lang: '',
    target_lang: '',
    source_text: '',
    target_text: '',
  });

  const postGlossaryTerm = async () => {
    try {
      console.log(data);
      const response = await addTermToGlossary(1, data);
      if (response.data) {
        console.log('post term res:', response.data);
        Message.success(t('glossary.import.success'));
      }
    } catch (error) {
      Message.error(t('glossary.import.fail'));
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
@/api/glossaries
