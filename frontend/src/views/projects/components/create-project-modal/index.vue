<template>
  <a-modal
    v-model:visible="visible"
    width="auto"
    :on-before-ok="handleSubmit"
    :title="$t('project.create')"
  >
    <a-form :model="form" :style="{ width: '600px' }">
      <a-form-item
        field="name"
        :label="$t('project.name')"
        :validate-trigger="['change', 'input']"
      >
        <a-input
          v-model="form.name"
          style="width: 380px"
          placeholder="Please enter the project name..."
        />
      </a-form-item>
      <a-form-item
        :label="$t('project.language')"
        :content-flex="false"
        :merge-props="false"
      >
        <a-row :gutter="8">
          <a-col :span="10">
            <a-form-item
              field="source language"
              validate-trigger="input"
              no-style
            >
              <a-select
                v-model="form.source_lang"
                :style="{ width: '180px' }"
                placeholder="source language"
              >
                <a-option>zh</a-option>
                <a-option>en</a-option>
                <a-option>ja</a-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :span="10">
            <a-form-item
              field="target language"
              validate-trigger="input"
              no-style
            >
              <a-select
                v-model="form.target_lang"
                :style="{ width: '180px' }"
                placeholder="target language"
              >
                <a-option>zh</a-option>
                <a-option>en</a-option>
                <a-option>ja</a-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>
      </a-form-item>
      <a-form-item
        field="source"
        :label="$t('project.source')"
        :validate-trigger="['change', 'input']"
      >
        <a-input
          v-model="form.comment"
          style="width: 380px"
          placeholder="Please enter the material source..."
        ></a-input>
      </a-form-item>
      <a-form-item field="deadline" :label="$t('project.deadline')">
        <a-date-picker
          v-model="form.deadline"
          placeholder="Please choose the deadline date..."
          style="width: 380px"
        />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
  import { ref, reactive, defineModel } from 'vue';
  import { createProject } from '@/api/projects';

  const visible = defineModel<boolean>('visible', {
    type: Boolean,
    default: false,
    required: true,
  });

  const form = reactive({
    name: '',
    source_lang: '',
    target_lang: '',
    deadline: '',
    comment: '',
  });

  const handleSubmit = async () => {
    form.deadline = new Date(form.deadline).toISOString();
    console.log(form);
    const response = await createProject(form);
    if (response.data) {
      console.log(response.data);
    }
  };
</script>

<style scoped></style>
