<template>
  <div ref="el" class="assign">
    <h4>Assign</h4>
    <el-input v-model="variable" @change="updateSelect" size="small" />
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
import nodeHeader from './nodeHeader.vue';
export default defineComponent({
  components: {
    nodeHeader,
  },
  setup() {
    const el = ref(null);
    const nodeId = ref(0);
    let df = null;
    const variable = ref(null);
    const dataNode = ref({});

    df = getCurrentInstance().appContext.config.globalProperties.$df.value;
    const updateSelect = (value) => {
      dataNode.value.data.variable = value;
      df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
    };

    onMounted(async () => {
      await nextTick();
      nodeId.value = el.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);

      variable.value = dataNode.value.data.variable;
    });

    return {
      el,
      variable,
      updateSelect,
    };
  },
});
</script>
<style scoped>
.assign {
  border-radius: 1rem;
  background-color: #f15300;
  padding: 1rem;
}
</style>
