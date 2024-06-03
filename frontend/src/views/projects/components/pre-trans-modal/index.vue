<template>
  <a-modal
    v-model:visible="visible"
    :on-before-ok="sendChoosenAPI"
    :title="$t('translate.pretrans')"
    width="auto"
  >
    <div>
      <a-radio-group class="custom-radio-group" v-model="choosenAPI">
        <a-radio :value="'deepl'">
          <template #radio="{ checked }">
            <a-space
              align="start"
              class="custom-radio-card"
              :class="{ 'custom-radio-card-checked': checked }"
            >
              <div className="custom-radio-card-mask">
                <div className="custom-radio-card-mask-dot" />
              </div>
              <div>
                <div className="custom-radio-card-title"> DeepL </div>
                <a-typography-text type="secondary">
                  A German machine translation service known for its
                  high-precision translations
                </a-typography-text>
              </div>
            </a-space>
          </template>
        </a-radio>
        <a-radio :value="'volc'">
          <template #radio="{ checked }">
            <a-space
              align="start"
              class="custom-radio-card"
              :class="{ 'custom-radio-card-checked': checked }"
            >
              <div className="custom-radio-card-mask">
                <div className="custom-radio-card-mask-dot" />
              </div>
              <div>
                <div className="custom-radio-card-title"> VolcEngine </div>
                <a-typography-text type="secondary">
                  A translation engine developed by Bytedance for a variety of
                  business scenarios
                </a-typography-text>
              </div>
            </a-space>
          </template>
        </a-radio>
      </a-radio-group>
    </div>
  </a-modal>
</template>

<script setup lang="ts">
  import { defineModel, ref } from 'vue';
  import { postTransAPI } from '@/api/machine';
  import { useRoute, useRouter } from 'vue-router';
  import { Message } from '@arco-design/web-vue';
  import { useI18n } from 'vue-i18n';

  const route = useRoute();
  const { t } = useI18n();
  const choosenAPI = ref('');
  const documentId = route.params.fileId;
  const visible = defineModel<boolean>('visible', {
    type: Boolean,
    default: false,
    required: true,
  });
  const emit = defineEmits(['pretransData']);

  const sendChoosenAPI = async () => {
    Message.loading({ content: t('translate.pretrans.loading'), duration: 0 });
    console.log('api:', choosenAPI.value, typeof choosenAPI.value);
    const data = {
      engine: choosenAPI.value,
      document_id: Number(documentId),
    };
    try {
      const response = await postTransAPI(data);
      if (response.data) {
        console.log(response);
        Message.success(t('translate.pretrans.success'));
        emit('pretransData', response.data);
      }
    } catch (error) {
      Message.error(t('translate.pretrans.fail'));
      console.error(error);
    } finally {
      Message.clear();
    }
  };
</script>

<style scoped>
  .custom-radio-group {
    display: flex;
    flex-wrap: nowrap;
    align-items: center;
    margin-top: 20px;
  }

  .custom-radio-card {
    flex: 1;
    padding: 10px 16px;
    border: 1px solid var(--color-border-2);
    border-radius: 4px;
    width: 300px;
    box-sizing: border-box;
  }

  .custom-radio-card-mask {
    height: 14px;
    width: 14px;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    border-radius: 100%;
    border: 1px solid var(--color-border-2);
    box-sizing: border-box;
  }

  .custom-radio-card-mask-dot {
    width: 8px;
    height: 8px;
    border-radius: 100%;
  }

  .custom-radio-card-title {
    color: var(--color-text-1);
    font-size: 14px;
    font-weight: bold;
    margin-bottom: 8px;
  }

  .custom-radio-card:hover,
  .custom-radio-card-checked,
  .custom-radio-card:hover .custom-radio-card-mask,
  .custom-radio-card-checked .custom-radio-card-mask {
    border-color: rgb(var(--primary-6));
  }

  .custom-radio-card-checked {
    background-color: var(--color-primary-light-1);
  }

  .custom-radio-card:hover .custom-radio-card-title,
  .custom-radio-card-checked .custom-radio-card-title {
    color: rgb(var(--primary-6));
  }

  .custom-radio-card-checked .custom-radio-card-mask-dot {
    background-color: rgb(var(--primary-6));
  }
</style>
