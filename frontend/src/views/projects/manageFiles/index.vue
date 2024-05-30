<template>
  <div
    class="container"
    :style="{ backgroundColor: theme === 'light' ? 'white' : 'black' }"
  >
    <div class="project-title">
      <a-typography-title :heading="4">{{ projectName }}</a-typography-title>
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
                <div class="icon-file"><icon-file /></div>
                <div
                  ><a-typography-text class="file-name">{{
                    $t('project.file.name')
                  }}</a-typography-text></div
                >
              </div>
              <div class="progress">
                <a-typography-text>{{
                  $t('project.file.progress')
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
              <div class="create_time">
                <a-typography-text>{{
                  $t('project.file.date')
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
          v-for="document in documents"
          :key="document.file_id"
          @dblclick="
            gotoTranslatePage(Number(route.params.projectId), document.id)
          "
        >
          <wrapper
            class="list-item"
            :style="{
              backgroundColor: theme === 'light' ? 'whitesmoke' : 'black',
            }"
          >
            <div class="left-side">
              <div class="filebox">
                <div class="icon-file"><icon-file /></div>
                <div class="file-name"
                  ><a-typography-text ellipsis>{{
                    document.name
                  }}</a-typography-text></div
                >
              </div>
              <div class="progress">
                <a-progress
                  :percent="0.6"
                  :style="{ width: '100%' }"
                ></a-progress>
              </div>
              <div class="type">
                <a-typography-text>{{ document.filetype }}</a-typography-text>
              </div>
              <div class="creater">
                <a-typography-text>{{ document.author }}</a-typography-text>
              </div>
            </div>
            <div class="right-side">
              <div class="create_time">
                <a-typography-text>{{
                  convertToBasicDateFormat(document.CreatedAt)
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
                    @ok="
                      handleDelete(
                        Number(route.params.projectId),
                        document.file_id
                      )
                    "
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
      <upload-file-modal v-model:visible="visible"></upload-file-modal>
    </main>
  </div>
</template>

<script setup lang="ts">
  import { computed, ref, onMounted } from 'vue';
  import { useRouter, useRoute } from 'vue-router';
  import { useAppStore } from '@/store';

  import UploadFileModal from '@/views/projects/components/upload-file-modal/index.vue';
  import { getProjectDetail } from '@/api/projects';
  import { deleteFileFromProject } from '@/api/files';

  const appStore = useAppStore();
  const router = useRouter();
  const route = useRoute();

  const visible = ref(false);
  const theme = computed(() => {
    return appStore.theme;
  });
  const documents = ref<any>([]);
  const projectName = ref<string>();

  const fetchFiles = async () => {
    try {
      const response = await getProjectDetail(Number(route.params.projectId));
      // console.log('response:', response);
      if (response && response.data) {
        documents.value = response.data.documents;
        projectName.value = response.data.project.name;
        console.log(documents);
      }
    } catch (error) {
      console.log('Fail to fetch files', error);
    }
  };

  const gotoTranslatePage = (projectId: number, fileId: number) => {
    router.push({
      name: 'translatePage',
      params: { projectId, fileId },
    });
  };

  const handleDelete = async (projectId: number, fileIds: number[]) => {
    try {
      await deleteFileFromProject(projectId, fileIds);
      fetchFiles();
    } catch (error) {
      console.error('Failed to delete document:', error);
    }
  };

  onMounted(fetchFiles);

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
      align-items: center;

      .icon-file {
        flex: 1;
      }

      .file-name {
        flex: 2;
        padding-top: 12px;
      }
    }

    .progress {
      flex: 2;
      display: flex;
      justify-content: center;
      padding: 0 10px;
      margin-left: 25px;
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
