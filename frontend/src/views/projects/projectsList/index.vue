<template>
  <div
    class="container"
    :style="{ backgroundColor: theme === 'light' ? 'white' : 'black' }"
  >
    <nav>
      <div class="create-project">
        <a-button type="primary" @click="visible = true">{{
          $t('menu.projects.createProject')
        }}</a-button>
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
                <a-typography-text class="project-name">{{
                  $t('project.name')
                }}</a-typography-text>
              </div>
              <div class="progress">
                <a-typography-text>{{
                  $t('project.progress')
                }}</a-typography-text>
              </div>
              <div class="language">
                <a-typography-text>{{
                  $t('project.language')
                }}</a-typography-text>
              </div>
              <div class="source">
                <a-typography-text>{{
                  $t('project.source')
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
          v-for="project in projects"
          :key="project.projectId"
          @click="gotoFileManage(project.projectId)"
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
                <a-typography-text class="project-name">{{
                  project.name
                }}</a-typography-text>
              </div>
              <div class="progress">
                <a-progress
                  :percent="0.6"
                  :style="{ width: '100%' }"
                ></a-progress>
              </div>
              <div class="language">
                <a-typography-text>{{ project.language }}</a-typography-text>
              </div>
              <div class="source">
                <a-typography-text>{{ project.source }}</a-typography-text>
              </div>
            </div>
            <div class="right-side">
              <div class="deadline">
                <a-typography-text>{{ project.deadline }}</a-typography-text>
              </div>
              <div class="icons">
                <div><icon-download /></div>
                <div class="delete"><icon-delete /></div>
              </div>
            </div>
          </wrapper>
        </a-list-item>
      </a-list>
      <create-project-modal v-model:visible="visible"></create-project-modal>
    </main>
  </div>
</template>

<script setup lang="ts">
  import { computed, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useAppStore } from '@/store';

  import CreateProjectModal from '@/views/projects/components/create-project-modal/index.vue';

  const appStore = useAppStore();
  const router = useRouter();

  const visible = ref(false);
  const theme = computed(() => {
    return appStore.theme;
  });

  const gotoFileManage = (projectId: any) => {
    console.log('enter funtion gotoFileManage');
    router.push({
      name: 'manageFiles',
      params: { projectId },
    });
  };

  const projects = [
    {
      projectId: 111,
      name: '项目1',
      progress: '',
      language: 'en -> zh',
      source: '光明出版社',
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

  nav {
    width: 82%;
    top: 140px;
    position: absolute;

    .create-project {
      margin-left: 140px;
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

      .project-name {
        flex: 2;
      }
    }

    .progress {
      flex: 2;
      display: flex;
      justify-content: center;
    }

    .language {
      flex: 1;
      display: flex;
      justify-content: center;
    }

    .source {
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
