<template>
  <div ref="variable_ref">
    <node title="Variable" color="#00a186">
      <el-select
        v-model="value"
        placeholder="Value"
        size="small"
        @change="updateVariableSelect"
        filterable
        allow-create
      >
        <el-option
          v-for="item in variableOptions"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
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
    const variable_ref = ref(null);
    const nodeId = ref(0);
    const value = ref(null);
    const dataNode = ref({});
    const df =
      getCurrentInstance().appContext.config.globalProperties.$df.value;
    const variableOptions = [];

    const updateVariableSelect = (value) => {
      updateNode(dataNode, nodeId, df, value, 'value');
    };

    onMounted(async () => {
      await nextTick();
      nodeId.value = variable_ref.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);
      getRootAssignNodes();
      value.value = dataNode.value.data.value;
    });

    const getRootAssignNodes = () => {
      const rootAssignNodes = df.getNodesFromName('Assign');
      rootAssignNodes.forEach((nodeId) => {
        const node = df.getNodeFromId(nodeId);
        if (
          node.data.variable &&
          node.outputs.output_1.connections.length === 0
        ) {
          variableOptions.push({
            value: node.data.variable,
            label: node.data.variable,
          });
        }
      });
    };

    return {
      variable_ref,
      value,
      updateVariableSelect,
      variableOptions,
    };
  },
});
</script>
