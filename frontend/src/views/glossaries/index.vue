<template>
  <div
    class="container"
    :style="{ backgroundColor: theme === 'light' ? 'white' : 'black' }"
  >
    <nav>
      <div class="import_grossary">
        <a-button type="primary" @click="visible = true">{{
          $t('grossary.import')
        }}</a-button>
      </div>
    </nav>
    <main> </main>
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

    .import_grossary {
      margin-left: 140px;
    }
  }

  main {
    margin-top: 100px;
    width: 82%;
  }
</style>
