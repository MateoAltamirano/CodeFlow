<template>
  <div ref="number_ref">
    <node title="Number" color="#00cd6a">
      <el-input-number
        v-model.number="value"
        @change="updateNumberInput"
        controls-position="right"
        size="small"
      />
    </node>
  </div>
</template>

<script>
import {
  defineComponent,
  onMounted,
  getCurrentInstance,
  ref,
  nextTick,
} from 'vue';
import node from './node.vue';
import { updateNode } from '../../utils/common';

export default defineComponent({
  components: {
    node,
  },

  setup() {
    const number_ref = ref(null);
    const nodeId = ref(0);
    const value = ref(null);
    const dataNode = ref({});
    const df =
      getCurrentInstance().appContext.config.globalProperties.$df.value;

    const updateNumberInput = (value) => {
      updateNode(dataNode, nodeId, df, value, 'value');
    };

    onMounted(async () => {
      await nextTick();
      nodeId.value = number_ref.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);
      value.value = dataNode.value.data.value;
    });

    return {
      number_ref,
      value,
      updateNumberInput,
    };
  },
});
</script>
