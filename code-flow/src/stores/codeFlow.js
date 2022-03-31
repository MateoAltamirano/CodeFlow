import { defineStore } from 'pinia';

export const useCodeFlowStore = defineStore({
  id: 'codeFlow',
  state: () => ({
    flow: {},
    code: '',
    result: '',
  }),
  actions: {
    updateFlow(flow) {
      this.flow = flow;
    },
    updateCode(code) {
      this.code = code;
    },
    updateResult(result) {
      this.result = result;
    },
  },
});
