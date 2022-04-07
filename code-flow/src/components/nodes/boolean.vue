<template>
  <div ref="boolean_ref">
    <node title="Boolean" color="#00cd6a">
      <el-select
        v-model="value"
        placeholder="Value"
        size="small"
        @change="updateBooleanSelect"
        filterable
        allow-create
      >
        <el-option
          v-for="item in selectOptions"
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
    const boolean_ref = ref(null);
    const nodeId = ref(0);
    const value = ref(null);
    const dataNode = ref({});
    const df =
      getCurrentInstance().appContext.config.globalProperties.$df.value;
    const selectOptions = [
      {
        value: 'True',
        label: 'True',
      },
      {
        value: 'False',
        label: 'False',
      },
    ];

    const updateBooleanSelect = (value) => {
      updateNode(dataNode, nodeId, df, value, 'value');
    };

    onMounted(async () => {
      await nextTick();
      nodeId.value = boolean_ref.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);
      value.value = dataNode.value.data.value;
    });

    return {
      boolean_ref,
      value,
      updateBooleanSelect,
      selectOptions,
    };
  },
});
</script>
