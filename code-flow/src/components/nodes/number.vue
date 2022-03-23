<template>
  <div ref="el" class="number">
    <h4>Number</h4>
    <el-input-number
      v-model.number="value"
      @change="updateSelect"
      controls-position="right"
      size="small"
    />
  </div>
</template>

<script>
import {
  defineComponent,
  onMounted,
  onBeforeUpdate,
  getCurrentInstance,
  readonly,
  ref,
  nextTick,
} from 'vue';
import nodeHeader from './nodeHeader.vue';
export default defineComponent({
  components: {
    nodeHeader,
  },
  setup() {
    const el = ref(null);
    const nodeId = ref(0);
    let df = null;
    const value = ref(null);
    const dataNode = ref({});

    df = getCurrentInstance().appContext.config.globalProperties.$df.value;
    const updateSelect = (value) => {
      dataNode.value.data.value = value;
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
.number {
  border-radius: 1rem;
  background-color: #00cd6a;
  padding: 1rem;
}
</style>
