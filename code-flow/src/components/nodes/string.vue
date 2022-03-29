<template>
  <div ref="el" class="assign">
    <h4>String</h4>
    <el-input v-model="value" @change="updateSelect" size="small" />
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

    df = getCurrentInstance().appContext.config.globalProperties.$df.value;
    const updateSelect = (value) => {
      dataNode.value.data.value = `"${value}"`;
      df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
    };

    onMounted(async () => {
      await nextTick();
      nodeId.value = el.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);

      value.value = dataNode.value.data.value;
    });

    return {
      el,
      value,
      updateSelect,
    };
  },
});
</script>
<style scoped>
.assign {
  border-radius: 1rem;
  background-color: #00cd6a;
  padding: 1rem;
}
</style>
