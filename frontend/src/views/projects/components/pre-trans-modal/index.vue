<template>
  <a-modal
    v-model:visible="visible"
    :on-before-ok="sendChoosenAPI"
    :title="$t('translate.pretrans')"
  >
    <div :style="{ marginTop: '20px' }">
      <a-radio-group>
        <template v-for="item in 2" :key="item">
          <a-radio :value="item">
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
                  <div className="custom-radio-card-title">
                    radio Card {{ item }}
                  </div>
                  <a-typography-text type="secondary">
                    this is a text
                  </a-typography-text>
                </div>
              </a-space>
            </template>
          </a-radio>
        </template>
      </a-radio-group>
    </div>
  </a-modal>
</template>

<script setup lang="ts">
  import { defineModel, ref } from 'vue';
  import { postTransAPI } from '@/api/machine';
  import { useRoute, useRouter } from 'vue-router';

  const route = useRoute()
  const choosenAPI = ref('');
  const document_id = route.params.fileId;
  const visible = defineModel<boolean>('visible', {
    type: Boolean,
    default: false,
    required: true,
  });

  const sendChoosenAPI = async() => {
    const data = {
      engine: choosenAPI,
      document_id: document_id
    }
    const response = await postTransAPI(choosenAPI.value, Number(document_id));
    if (response.data) {
      console.log(response.data);
    }
  };
</script>

<style scoped>
  .custom-radio-card {
    padding: 10px 16px;
    border: 1px solid var(--color-border-2);
    border-radius: 4px;
    width: 250px;
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
