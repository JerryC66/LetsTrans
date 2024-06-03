<template>
  <div
    class="container"
    :style="{ backgroundColor: theme === 'light' ? 'white' : 'black' }"
  >
    <nav>
      <div class="row1">
        <div class="icons-bar">
          <a-space size="medium">
            <icon-oblique-line size="24" />
            <icon-copy size="24" />
            <icon-quote size="24" />
            <icon-redo size="24" />
            <icon-scissor size="24" />
            <icon-undo size="24" />
            <icon-zoom-in size="24" />
            <icon-zoom-out size="24" />
            <icon-sort size="24" />
            <icon-underline size="24" />
          </a-space>
        </div>
        <div class="record-bar">
          <div class="file-name">
            <a-typography-text>{{ document?.name }}</a-typography-text>
          </div>
          <div class="language">
            <a-typography-text
              >{{ document?.source_lang }} >
              {{ document?.target_lang }}</a-typography-text
            >
          </div>
          <div class="progress">
            <a-progress
              :percent="document?.progress"
              :style="{ width: '100%' }"
            ></a-progress>
          </div>
        </div>
      </div>
      <div class="row2">
        <div class="search-bar">
          <a-input-search
            size="large"
            :style="{ width: '740px' }"
            placeholder="Please enter something"
            search-button
          />
        </div>
        <div class="btn-group">
          <div class="download">
            <a-button type="primary" size="large">
              <template #icon>
                <icon-download size="large" />
              </template>
            </a-button>
          </div>
          <div class="import">
            <a-button
              @click="visible1 = true"
              size="large"
              type="outline"
              shape="round"
              >{{ $t('translate.addterm') }}</a-button
            >
          </div>
          <div class="pre-trans">
            <a-button
              @click="visible2 = true"
              size="large"
              type="outline"
              shape="round"
              >{{ $t('translate.pretrans') }}</a-button
            >
          </div>
        </div>
      </div>
    </nav>
    <main>
      <div class="left-side">
        <div
          class="block"
          v-for="(segment, index) in segments"
          :key="'block-' + index"
          
        >
          <div class="left-wrapper">
            <div class="left-block" :key="'left-' + index">
              <a-textarea
                v-model="segment.source_text"
                :auto-size="{
                  minRows: rowsArray[index],
                }"
              ></a-textarea>
            </div>
          </div>
          <div class="right-wrapper">
            <div class="right-block" :key="'right-' + index">
              <a-textarea
                v-model="pretransRes[index].target_text"
                :auto-size="{
                  minRows: rowsArray[index],
                }"
                @focus="handleFocus(index)"
                @blur="handleBlur"
              ></a-textarea>
            </div>
          </div>
          <div class="confirm-btn">
            <a-button
              type="primary"
              size="large"
              :disabled="activeBlock !== index"
              @click="updateSegment(index)"
            >
              <template #icon>
                <icon-check size="large" />
              </template>
            </a-button>
          </div>
        </div>
      </div>

      <div class="right-side">
        <a-list>
          <a-list-item>hello</a-list-item>
        </a-list>
      </div>
    </main>
  </div>
  <import-grossary-modal v-model:visible="visible1"></import-grossary-modal>
  <pre-trans-modal
    v-model:visible="visible2"
    @pretransData="handlePretrans"
  ></pre-trans-modal>
</template>

<script setup lang="ts">
  import { computed, ref, reactive, onMounted, watch, nextTick } from 'vue';
  import { useRoute, useRouter } from 'vue-router';
  import { useAppStore } from '@/store';
  import PreTransModal from '@/views/projects/components/pre-trans-modal/index.vue';
  import ImportGrossaryModal from '@/views/projects/components/import-grossary-modal/index.vue';
  import { getDocumentSegments, updateDocumentSegment } from '@/api/translate';

  const appStore = useAppStore();
  const router = useRouter();
  const route = useRoute();

  const visible1 = ref(false);
  const visible2 = ref(false);
  const theme = computed(() => {
    return appStore.theme;
  });

  const segments = ref<any>([]);
  const document = ref<any>();
  const pretransRes = ref<any>([]);
  const rowsArray = ref<any>([]);
  const activeBlock = ref(-1);

  const calculateRows = (text1: string, text2: string) => {
    const isChinese = (text) => /[\u4e00-\u9fa5]/.test(text);
    const getRows = (text, isChineseText) => {
      if (!text) return 2;
      const charsPerLine = isChineseText ? 20 : 40;
      return Math.max(3, Math.ceil(text.length / charsPerLine));
    };

    const chinese1 = isChinese(text1);
    const chinese2 = isChinese(text2);
    const row1 = getRows(text1, chinese1);
    const row2 = getRows(text2, chinese2);
    const maxRows = Math.max(row1, row2);

    return maxRows;
  };

  const updateRows = () => {
    const newRows = segments.value.map((segment, index) => {
      const pretransText = pretransRes.value[index]
        ? pretransRes.value[index].target_text
        : '';
      return calculateRows(segment.source_text, pretransText);
    });
    rowsArray.value = newRows;
  };

  const fetchSegments = async () => {
    try {
      const response = await getDocumentSegments(Number(route.params.fileId));
      console.log('response:', response);
      if (response && response.data) {
        segments.value = response.data.segments;
        document.value = response.data.document;
        pretransRes.value = segments.value.map(() => ({ target_text: '' }));
        updateRows();
      }
    } catch (error) {
      console.log('Fail to fetch segments', error);
    }
  };

  const handlePretrans = (data) => {
    pretransRes.value = data.segments;
    updateRows();
  };

  const updateSegment = async (index) => {
    const segment = segments.value[index];
    const newText = pretransRes.value[index].target_text;
    const documentId = Number(route.params.fileId);
    const segmentId = segments.value[index].id;
    console.log(newText, documentId, segmentId);
    if (segment && newText) {
      const data = {
        targetText: newText,
        finished: true,
      };
      try {
        console.log('尝试请求');
        const response = await updateDocumentSegment(
          documentId,
          segmentId,
          data
        );
        if (response && response.data) {
          console.log('Succeed to update segment', response.data);
        }
      } catch (error) {
        console.log('Fail to update segment', error);
      }
    }
    updateRows();
  };

const handleFocus = (index) => {
  activeBlock.value = index;
};

const handleBlur = async () => {
  await nextTick(); 
  setTimeout(() => {
    activeBlock.value = -1;
  }, 100); 
};

  watch([segments, pretransRes], updateRows, { deep: true });

  onMounted(fetchSegments);
</script>

<style scoped>
  .container {
    margin: 28px 28px;
    height: 92vh;
    margin-bottom: 0;
    display: flex;
    flex-direction: column;
    align-items: center;
    overflow: hidden;
  }

  nav {
    width: 80%;
    top: 120px;
    position: absolute;
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;

    .row1 {
      display: flex;

      .icons-bar {
        flex: 1;
        margin-right: 100px;
        margin-left: 50px;
      }

      .record-bar {
        flex: 2;
        display: flex;

        .file-name {
          flex: 2;
        }

        .language {
          flex: 2;
        }

        .progress {
          flex: 3;
          margin-right: 100px;
        }
      }
    }

    .row2 {
      display: flex;
      margin-top: 30px;

      .search-bar {
        flex: 3;
        display: flex;
        justify-content: center;
      }

      .btn-group {
        flex: 1;
        display: flex;
        justify-content: space-evenly;
        padding-right: 20px;

        .download {
          width: 36px;
          height: 36px;
          border-radius: 30%;
          background-color: #165dff;
          color: white;
          display: flex;
          justify-content: center;
          align-items: center;
        }
      }
    }
  }

  main {
    position: absolute;
    top: 150px;
    margin-top: 100px;
    padding-bottom: 20px;
    width: 80%;
    display: flex;
    height: calc(100vh - 200px);

    .left-side {
      flex: 3;
    }

    .right-side {
      flex: 1;
      padding-right: 20px;
    }

    .left-side .block {
      display: flex;
      margin-left: 50px;
      margin-right: 50px;

      .left-wrapper {
        flex: 1;
        margin-right: 20px;

        .left-block {
          margin-bottom: 20px;
        }
      }

      .right-wrapper {
        flex: 1;

        .right-block {
          margin-bottom: 20px;
        }
      }

      .confirm-btn {
        margin-left: 25px;
      }
    }
  }
  main .left-side {
    overflow-y: scroll;
  }

  main .left-side::-webkit-scrollbar {
    display: none;
  }
</style>
