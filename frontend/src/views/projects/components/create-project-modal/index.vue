<template>
  <a-modal v-model:visible="visible" width="auto" :title="$t('project.create')">
    <a-form :model="form" :style="{ width: '600px' }" @submit="handleSubmit">
      <a-form-item
        field="name"
        :label="$t('project.name')"
        :rules="[{ required: true, message: 'project name is required' }]"
        :validate-trigger="['change', 'input']"
      >
        <a-input
          v-model="form.project_name"
          style="width: 380px"
          placeholder="please enter the project name..."
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
              :rules="[{ required: true, message: 'firstname is required' }]"
              no-style
            >
              <a-select
                v-modal="form.source_language"
                :style="{ width: '180px' }"
                placeholder="source language"
              >
                <a-option>zh</a-option>
                <a-option>en</a-option>
                <a-option>jp</a-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :span="10">
            <a-form-item
              field="target language"
              validate-trigger="input"
              :rules="[{ required: true, message: 'lastname is required' }]"
              no-style
            >
              <a-select
                v-modal="form.target_language"
                :style="{ width: '180px' }"
                placeholder="target language"
              >
                <a-option>zh</a-option>
                <a-option>en</a-option>
                <a-option>jp</a-option>
              </a-select>
            </a-form-item>
          </a-col>
        </a-row>
      </a-form-item>
      <a-form-item
        field="source"
        :label="$t('project.source')"
        :rules="[{ required: true, message: 'project name is required' }]"
        :validate-trigger="['change', 'input']"
      >
        <a-input
          v-model="form.source"
          style="width: 380px"
          placeholder="Please enter the material source..."
        ></a-input>
      </a-form-item>
      <a-form-item field="deadline" :label="$t('project.deadline')">
        <a-date-picker
          placeholder="Please choose the deadline date..."
          style="width: 380px"
        />
      </a-form-item>
      <a-form-item class="btn-container">
        <a-button class="submit-btn" type="primary"> Submit</a-button>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
  import { ref, reactive, defineModel } from 'vue';

  const visible = defineModel<boolean>('visible', {
    type: Boolean,
    default: false,
    required: true,
  });

  const form = reactive({
    project_name: '',
    source_language: '',
    target_language: '',
    source: '',
    deadline: '',
  });
  const handleSubmit = () => {};
</script>

<style scoped></style>
