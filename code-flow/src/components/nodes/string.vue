<template>
  <div ref="string_ref">
    <node title="String" color="#00cd6a">
      <el-input v-model="value" @change="updateStringInput" size="small" />
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
    const string_ref = ref(null);
    const nodeId = ref(0);
    let df = null;
    const value = ref(null);
    const dataNode = ref({});

    df = getCurrentInstance().appContext.config.globalProperties.$df.value;
    const updateStringInput = (value) => {
      updateNode(dataNode, nodeId, df, `"${value}"`, 'value');
    };

    onMounted(async () => {
      await nextTick();
      nodeId.value = string_ref.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);
      const valueWithoutQuotes = dataNode.value.data.value.replaceAll('"', '');
      value.value = valueWithoutQuotes;
    });

    return {
      string_ref,
      value,
      updateStringInput,
    };
  },
});
</script>
