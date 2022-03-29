<template>
  <div ref="el" class="add">
    <h4>If</h4>
    <el-select
      v-model="operator"
      class="m-2"
      placeholder="Operator"
      size="small"
      @change="updateSelect"
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
  onBeforeUpdate,
  onUpdated,
  getCurrentInstance,
  readonly,
  ref,
  nextTick,
} from 'vue';
export default defineComponent({
  setup() {
    const el = ref(null);
    const nodeId = ref(0);
    let df = null;
    const operator = ref(null);
    const dataNode = ref({});

    df = getCurrentInstance().appContext.config.globalProperties.$df.value;
    const updateSelect = (value) => {
      dataNode.value.data.operator = value;
      df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
    };

    onMounted(async () => {
      await nextTick();
      nodeId.value = el.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);

      operator.value = dataNode.value.data.operator;
    });

    const options = [
      {
        value: '==',
        label: '==',
      },
      {
        value: '!=',
        label: '!=',
      },
      {
        value: '>',
        label: '>',
      },
      {
        value: '<',
        label: '<',
      },
      {
        value: '>=',
        label: '>=',
      },
      {
        value: '<=',
        label: '<=',
      },
    ];
    return {
      el,
      options,
      operator,
      updateSelect,
    };
  },
});
</script>
<style scoped>
.add {
  border-radius: 1rem;
  background-color: #0099de;
  padding: 1rem;
}
</style>
