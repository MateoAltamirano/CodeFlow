<template>
  <div ref="assign_ref">
    <node title="Assign" color="#f15300">
      <el-input v-model="variable" @change="updateAssignInput" size="small" />
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
    const assign_ref = ref(null);
    const nodeId = ref(0);
    const variable = ref(null);
    const dataNode = ref({});
    const df =
      getCurrentInstance().appContext.config.globalProperties.$df.value;

    const updateAssignInput = (value) => {
      updateNode(dataNode, nodeId, df, value, 'variable');
    };

    onMounted(async () => {
      await nextTick();
      nodeId.value = assign_ref.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);
      variable.value = dataNode.value.data.variable;
    });

    return {
      assign_ref,
      variable,
      updateAssignInput,
    };
  },
});
</script>
