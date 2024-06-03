<template>
  <div
    class="container"
    :style="{ backgroundColor: theme === 'light' ? 'white' : 'black' }"
  >
    <nav>
      <div class="import_glossary">
        <a-button type="primary" size="large" @click="visible = true">{{
          $t('glossary.import')
        }}</a-button>
      </div>
    </nav>
    <main>
      <div class="grossaries-cards">
        <a-row :gutter="16">
          <a-col
            :span="8"
            v-for="glossary in glossariesRes.glossaries"
            :key="glossary.id"
          >
            <a-card :title="glossary.name" :bordered="true">
              <template #extra>
                <a-link @click="handleMore(glossary.id)">More</a-link>
              </template>
              <template #cover>
                <div
                  :style="{
                    height: '180px',
                    overflow: 'hidden',
                  }"
                >
                  <img
                    :style="{ width: '100%', transform: 'translateY(-20px)' }"
                    alt="dessert"
                    src="https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/a20012a2d4d5b9db43dfc6a01fe508c0.png~tplv-uwbnlip3yd-webp.webp"
                  />
                </div>
              </template>
              <a-card-meta>
                <template #description>
                  <div style="line-height: 1.2; margin-bottom: 6px">
                    Author: {{ glossary.author }}
                  </div>
                  <div style="line-height: 1.2">
                    Language: {{ glossary.source_lang }} >
                    {{ glossary.target_lang }}
                  </div>
                </template>
              </a-card-meta>
            </a-card>
          </a-col>
        </a-row>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
  import { computed, ref, onMounted } from 'vue';
  import { useRouter } from 'vue-router';
  import { useAppStore } from '@/store';

  import { getGlossaries } from '@/api/grossaries';

  const appStore = useAppStore();
  const router = useRouter();
  const visible = ref(false);
  const theme = computed(() => {
    return appStore.theme;
  });
  const glossariesRes = ref<any>([]);

  const fetchGlossaries = async () => {
    try {
      const response = await getGlossaries();
      if (response && response.data) {
        glossariesRes.value = response.data;
        console.log(glossariesRes);
      }
    } catch (error) {
      console.error('Failed to fetch grossaries:', error);
    }
  };

  const handleMore = (glossaryId: number) => {};

  onMounted(fetchGlossaries);
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

    .import_glossary {
      margin-left: 120px;
    }
  }

  main {
    margin-top: 100px;
    width: 82%;
    height: calc(100vh - 200px);
    padding-top: 20px;
  }
</style>
