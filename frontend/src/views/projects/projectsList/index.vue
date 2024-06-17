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
                <a-typography-text>{{ $t('project.name') }}</a-typography-text>
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
                <div class="download"><icon-download /></div>
                <div class="delete"><icon-delete /></div>
              </div>
            </div>
          </wrapper>
        </template>
        <a-list-item
          v-for="project in projects"
          :key="project.id"
          @dblclick="gotoFileManage(project.id)"
        >
          <wrapper
            class="list-item"
            :style="{
              backgroundColor: theme === 'light' ? 'whitesmoke' : 'black',
            }"
          >
            <div class="left-side">
              <div class="filebox flex">
                <icon-file class="icon-file" />
                <div class="project-name">
                  <a-typography-text ellipsis>{{
                    project.name
                  }}</a-typography-text>
                </div>
              </div>
              <div class="progress">
                <a-progress
                  :percent="project.progress"
                  :style="{ width: '100%' }"
                ></a-progress>
              </div>
              <div class="language">
                <a-typography-text ellipsis
                  >{{ project.source_lang }} >
                  {{ project.target_lang }}</a-typography-text
                >
              </div>
              <div class="source">
                <a-typography-text ellipsis>{{
                  project.comment
                }}</a-typography-text>
              </div>
            </div>
            <div class="right-side">
              <div class="deadline">
                <a-typography-text>{{
                  convertToBasicDateFormat(project.deadline)
                }}</a-typography-text>
              </div>
              <div class="icons">
                <div class="download">
                  <a-popconfirm content="conform to download this project?">
                    <icon-download />
                  </a-popconfirm>
                </div>
                <div class="delete">
                  <a-popconfirm
                    @ok="handleDelete([project.id])"
                    content="conform to delete this project?"
                    type="error"
                  >
                    <icon-delete />
                  </a-popconfirm>
                </div>
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
  import { computed, ref, onMounted } from 'vue';
  import { useRouter } from 'vue-router';
  import { useAppStore } from '@/store';
  import CreateProjectModal from '@/views/projects/components/create-project-modal/index.vue';
  import { getProjects, deleteProject } from '@/api/projects';

  const appStore = useAppStore();
  const router = useRouter();
  const visible = ref(false);
  const theme = computed(() => {
    return appStore.theme;
  });
  const projects = ref<any>([]);

  const fetchProjects = async () => {
    try {
      const response = await getProjects();
      if (response && response.data) {
        projects.value = response.data;
        console.log(projects);
      }
    } catch (error) {
      console.error('Failed to fetch projects:', error);
    }
  };

  const handleDelete = async (projectId: number[]) => {
    try {
      await deleteProject(projectId);
      fetchProjects();
    } catch (error) {
      console.error('Failed to delete project:', error);
    }
  };

  onMounted(fetchProjects);

  const gotoFileManage = (projectId: any) => {
    router.push({
      name: 'manageFiles',
      params: { projectId },
    });
  };

  const convertToBasicDateFormat = (isoDateString) => {
    return isoDateString.split('T')[0];
  };
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
    height: calc(100vh - 200px);
    overflow-y: scroll;
  }

  .list {
    background-color: rgb(255, 252, 252);
    /* overflow-y: scroll; */
  }

  main::-webkit-scrollbar {
    display: none;
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
      align-items: center;

      .icon-file {
        flex: 1;
      }

      .project-name {
        flex: 2;
        padding-top: 14px;
      }
    }

    .progress {
      flex: 2;
      display: flex;
      justify-content: center;
      margin: 0 20px 0 25px;
    }

    .language {
      flex: 1;
      display: flex;
      justify-content: center;
      margin: 0 5px;
      padding-top: 14px;
    }

    .source {
      flex: 1;
      padding: 0 5px;
      padding-top: 14px;
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
