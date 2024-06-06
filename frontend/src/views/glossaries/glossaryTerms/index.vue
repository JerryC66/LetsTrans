<template>
  <div
    class="container"
    :style="{ backgroundColor: theme === 'light' ? 'white' : 'black' }"
  >
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
            <icon-edit />
            <icon-delete />
          </template>
        </a-list-item>
      </a-list>
    </div>
  </div>
</template>

<script setup lang="ts">
  import { computed, ref, onMounted } from 'vue';
  import { useRouter, useRoute } from 'vue-router';
  import { useAppStore } from '@/store';

  import { getGlossaryTerms } from '@/api/glossaries';

  const appStore = useAppStore();
  const router = useRouter();
  const route = useRoute();

  const visible = ref(false);
  const theme = computed(() => {
    return appStore.theme;
  });
  const glossaryId = Number(route.params.glossaryId);

  interface Term {
    id: number;
    document_id: string;
    created_at: string;
    updated_at: string;
    source_lang: string;
    target_lang: string;
    source_text: string;
    target_text: string;
  }
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

  .terms-list {
    margin-top: 80px;
    overflow-y: scroll;
  }

  .terms-list::-webkit-scrollbar {
    display: none;
  }

  .wrapper {
    display: flex;
    justify-content: space-between;
  }
</style>
