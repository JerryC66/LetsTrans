import { defineStore } from 'pinia';

const useTranslationStore = defineStore('translation', {
  state: () => ({
    documents: {},
  }),
  actions: {
    addDocument(id, documentData) {
      if (!this.documents[id]) {
        this.documents[id] = {
          segments: documentData.segments,
          name: documentData.document.name,
          source_lang: documentData.document.source_lang,
          target_lang: documentData.document.target_lang,
          progress: documentData.document.progress,
        };
      }
    },
    updateSegment(documentId, segmentIndex, targetText) {
      const document = this.documents[documentId];
      if (document && document.segments[segmentIndex]) {
        document.segments[segmentIndex].target_text = targetText;
      }
    },
  },
});

export default useTranslationStore;
