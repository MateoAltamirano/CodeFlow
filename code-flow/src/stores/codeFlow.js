import { defineStore } from 'pinia';

export const useCodeFlowStore = defineStore({
  id: 'codeFlow',
  state: () => ({
    flow: {},
    code: '',
    result: '',
  }),
  actions: {
    updateStore(key, value) {
      switch (key) {
        case 'flow':
          this.flow = value;
          break;
        case 'code':
          this.code = value;
          break;
        case 'result':
          this.result = value;
          break;
        default:
          break;
      }
    },
  },
});
