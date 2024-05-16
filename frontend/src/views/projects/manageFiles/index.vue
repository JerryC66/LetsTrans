<template>
  <div
    class="container"
    :style="{ backgroundColor: theme === 'light' ? 'white' : 'black' }"
  >
    <div class="project-title">
      <a-typography-title :heading="4">Project X</a-typography-title>
    </div>

    <nav>
      <div class="button-groups">
        <div>
          <a-button type="primary" status="success" @click="visible = true">{{
            $t('project.file.upload')
          }}</a-button>
        </div>
        <div>
          <a-button type="primary">{{ $t('project.files.download') }}</a-button>
        </div>
        <div>
          <a-button type="primary" status="danger">{{
            $t('project.files.delete')
          }}</a-button>
        </div>
      </div>
    </nav>
    <main>
      <a-list
        class="list"
        size="small"
        :split="false"
        :bordered="false"
        :style="{ backgroundColor: theme === 'light' ? 'white' : 'black' }"
      >
        <template #header>
          <wrapper
            class="list-item"
            :style="{
              backgroundColor: theme === 'light' ? 'whitesmoke' : 'black',
            }"
          >
            <div class="left-side">
              <div class="filebox">
                <icon-file class="icon-file" />
                <a-typography-text class="file-name">{{
                  $t('project.file.name')
                }}</a-typography-text>
              </div>
              <div class="progress">
                <a-typography-text>{{
                  $t('project.progress')
                }}</a-typography-text>
              </div>
              <div class="type">
                <a-typography-text>{{
                  $t('project.file.type')
                }}</a-typography-text>
              </div>
              <div class="creater">
                <a-typography-text>{{
                  $t('project.file.creater')
                }}</a-typography-text>
              </div>
            </div>
            <div class="right-side">
              <div class="deadline">
                <a-typography-text>{{
                  $t('project.deadline')
                }}</a-typography-text>
              </div>
              <div class="icons">
                <div><icon-download /></div>
                <div class="delete"><icon-delete /></div>
              </div>
            </div>
          </wrapper>
        </template>
        <a-list-item
          v-for="file in files"
          :key="file.fileId"
          @dblclick="gotoTranslatePage(route.params.projectId, file.fileId)"
        >
          <wrapper
            class="list-item"
            :style="{
              backgroundColor: theme === 'light' ? 'whitesmoke' : 'black',
            }"
          >
            <div class="left-side">
              <div class="filebox">
                <icon-file class="icon-file" />
                <a-typography-text class="file-name">{{
                  file.name
                }}</a-typography-text>
              </div>
              <div class="progress">
                <a-progress
                  :percent="0.6"
                  :style="{ width: '100%' }"
                ></a-progress>
              </div>
              <div class="type">
                <a-typography-text>{{ file.type }}</a-typography-text>
              </div>
              <div class="creater">
                <a-typography-text>{{ file.creater }}</a-typography-text>
              </div>
            </div>
            <div class="right-side">
              <div class="deadline">
                <a-typography-text>{{ file.deadline }}</a-typography-text>
              </div>
              <div class="icons">
                <div><icon-download /></div>
                <div class="delete"><icon-delete /></div>
              </div>
            </div>
          </wrapper>
        </a-list-item>
      </a-list>
      <upload-file-modal v-model:visible="visible"></upload-file-modal>
    </main>
  </div>
</template>

<script setup lang="ts">
  import { computed, ref } from 'vue';
  import { useRouter, useRoute } from 'vue-router';
  import { useAppStore } from '@/store';

  import UploadFileModal from '@/views/projects/components/upload-file-modal/index.vue';

  const appStore = useAppStore();
  const router = useRouter();
  const route = useRoute();

  const visible = ref(false);
  const theme = computed(() => {
    return appStore.theme;
  });

  const gotoTranslatePage = (projectId: any, fileId: any) => {
    router.push({
      name: 'translatePage',
      params: { projectId, fileId },
    });
  };

  const files = [
    {
      fileId: 111,
      name: '文件1',
      progress: '',
      type: 'doc',
      creater: 'cc',
      deadline: '2024.6.1',
    },
  ];
</script>

<style scoped>
  .container {
    margin: 28px 28px;
    height: 92%;
    margin-bottom: 0;
    display: flex;
    flex-direction: column;
    align-items: center;
  }

  .project-title {
    position: absolute;
    left: 350px;
    top: 110px;
    font-weight: 800;
  }

  nav {
    width: 82%;
    top: 140px;
    position: absolute;
    display: flex;
    justify-content: flex-end;

    .button-groups {
      margin-right: 10%;
      display: flex;
      justify-content: space-evenly;

      div {
        margin-left: 15px;
      }
    }
  }

  main {
    margin-top: 100px;
    width: 82%;
  }

  .list {
    background-color: rgba(252, 252, 252, 0.755);
  }

  .list-item {
    display: flex;
    align-items: center;
    height: 40px;
    padding-right: 20px;
  }

  .left-side {
    flex: 5;
    display: flex;
    align-items: center;

    .filebox {
      flex: 1;
      display: flex;

      .icon-file {
        flex: 1;
      }

      .file-name {
        flex: 2;
      }
    }

    .progress {
      flex: 2;
      display: flex;
      justify-content: center;
    }

    .type {
      flex: 1;
      display: flex;
      justify-content: center;
    }

    .creater {
      flex: 1;
    }
  }

  .right-side {
    flex: 3;
    display: flex;
    align-items: center;
    justify-content: flex-end;

    .icons {
      display: flex;
      margin-left: 50px;
      justify-content: center;

      .delete {
        margin-left: 12px;
      }
    }
  }
</style>
