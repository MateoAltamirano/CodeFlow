<template>
  <div ref="el" class="assign">
    <h4>For</h4>
    <el-input v-model="variable" @change="updateSelect" size="small" />
    <el-input-number
      v-model.number="start"
      @change="updateStart"
      controls-position="right"
      size="small"
    />
    <el-input-number
      v-model.number="stop"
      @change="updateStop"
      controls-position="right"
      size="small"
    />
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
    const variable = ref(null);
    const start = ref(null);
    const stop = ref(null);
    const dataNode = ref({});

    df = getCurrentInstance().appContext.config.globalProperties.$df.value;
    const updateSelect = (value) => {
      dataNode.value.data.variable = value;
      df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
    };
    const updateStart = (value) => {
      dataNode.value.data.start = value;
      df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
    };
    const updateStop = (value) => {
      dataNode.value.data.stop = value;
      df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
    };

    onMounted(async () => {
      await nextTick();
      nodeId.value = el.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);

      variable.value = dataNode.value.data.variable;
      start.value = dataNode.value.data.start;
      stop.value = dataNode.value.data.stop;
    });

    return {
      el,
      variable,
      start,
      stop,
      updateSelect,
      updateStart,
      updateStop,
    };
  },
});
</script>
<style scoped>
.assign {
  border-radius: 1rem;
  background-color: #ae59b9;
  padding: 1rem;
}
</style>
