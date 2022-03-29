<template>
  <div ref="el" class="assign">
    <h4>Variable</h4>
    <el-select
      v-model="value"
      placeholder="Value"
      size="small"
      @change="updateSelect"
      filterable
      allow-create
    >
      <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      />
    </el-select>
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
export default defineComponent({
  setup() {
    const el = ref(null);
    const nodeId = ref(0);
    let df = null;
    const value = ref(null);
    const dataNode = ref({});
    const options = [];

    df = getCurrentInstance().appContext.config.globalProperties.$df.value;
    const updateSelect = (value) => {
      dataNode.value.data.value = value;
      df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
    };

    onMounted(async () => {
      await nextTick();
      nodeId.value = el.value.parentElement.parentElement.id.slice(5);
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
          options.push({
            value: node.data.variable,
            label: node.data.variable,
          });
        }
      });
    };

    return {
      el,
      value,
      updateSelect,
      options,
    };
  },
});
</script>
<style scoped>
.assign {
  border-radius: 1rem;
  background-color: #00a186;
  padding: 1rem;
}
</style>
