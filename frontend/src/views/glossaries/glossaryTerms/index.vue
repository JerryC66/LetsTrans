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
    <div class="terms-list">
      <a-list>
        <a-list-item v-for="(term, index) in terms" :key="index">
          <div class="wrapper">
            <div class="term-text">
              {{ term.source_text }}
              <a-divider direction="vertical"></a-divider>
              {{ term.target_text }}
            </div>
            <div class="term-language">
              {{ term.source_lang }}
              >
              {{ term.target_lang }}
            </div>
          </div>

          <template #actions>
            <icon-edit @click="visible = true" />
            <icon-delete />
          </template>
        </a-list-item>
      </a-list>
    </div>
  </div>
  <edit-term-modal v-model:visible="visible"></edit-term-modal>
</template>

<script setup lang="ts">
  import { computed, ref, onMounted } from 'vue';
  import { useRouter, useRoute } from 'vue-router';
  import { useAppStore } from '@/store';

  import editTermModal from '@/views/glossaries/components/edit-term-modal/index.vue';

  import { getGlossaryTerms } from '@/api/glossaries';

  const appStore = useAppStore();
  const router = useRouter();
  const route = useRoute();

  const visible = ref(false);
  const theme = computed(() => {
    return appStore.theme;
  });
  const glossaryId = Number(route.params.glossaryId);

  const terms = ref<any>([]);
  const glossaryName = ref<string>('');

  const fetchTerms = async () => {
    try {
      const response = await getGlossaryTerms(glossaryId);
      if (response && response.data) {
        console.log('terms response:', response.data);
        glossaryName.value = response.data.glossary;
        terms.value = response.data.terms;
      }
    } catch (error) {
      console.log('Fail to fetch files', error);
    }
  };

  onMounted(fetchTerms);
</script>

<style scoped>
  .container {
    margin: 28px 28px;
    height: 92vh;
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
      margin-left: 180px;
    }
  }

  .terms-list {
    width: 70%;
    margin-top: 120px;
    overflow-y: scroll;
  }

  .terms-list::-webkit-scrollbar {
    display: none;
  }

  .wrapper {
    display: flex;
    justify-content: space-between;
    margin-right: 100px;
    padding-left: 20px;
  }
</style>
