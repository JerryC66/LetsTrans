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
            <a-typography-text>文件1.docx</a-typography-text>
          </div>
          <div class="language">
            <a-typography-text>简体中文</a-typography-text>
          </div>
          <div class="progress">
            <a-progress :percent="0.6" :style="{ width: '100%' }"></a-progress>
          </div>
          <div class="description">
            <a-typography-text> 6000字/共10000字 </a-typography-text>
          </div>
        </div>
      </div>
      <div class="row2">
        <div class="search-bar">
          <a-input-search
            size="large"
            :style="{ width: '750px' }"
            placeholder="Please enter something"
            search-button
          />
        </div>
        <div class="btn-group">
          <div class="download">
            <icon-download size="18" />
          </div>
          <div class="import">
            <a-button size="large" type="primary" shape="round">{{
              $t('translate.addterm')
            }}</a-button>
          </div>
          <div class="pre-trans">
            <a-button size="large" type="primary" shape="round">{{
              $t('translate.pretrans')
            }}</a-button>
          </div>
        </div>
      </div>
    </nav>
    <main>
      <div class="left-side">
        <div class="block">
          <div class="left-block">
            <a-textarea> </a-textarea>
          </div>
          <div class="right-block">
            <a-textarea></a-textarea>
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
</template>

<script setup lang="ts">
  import { computed } from 'vue';
  import { useAppStore } from '@/store';

  const appStore = useAppStore();

  const theme = computed(() => {
    return appStore.theme;
  });
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
          flex: 1;
        }
        .language {
          flex: 1;
        }
        .progress {
          flex: 2;
        }
        .description {
          flex: 1.5;
          margin-left: 50px;
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
    width: 80%;
    display: flex;
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
      .left-block {
        flex: 1;
        margin-right: 20px;
      }
      .right-block {
        flex: 1;
      }
    }
  }
</style>
