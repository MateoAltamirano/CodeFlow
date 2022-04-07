<template>
  <div ref="if_ref">
    <node title="If" color="#0099de">
      <el-select
        v-model="operator"
        class="m-2"
        placeholder="Operator"
        size="small"
        @change="updateIfSelect"
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
    const if_ref = ref(null);
    const nodeId = ref(0);
    const operator = ref(null);
    const dataNode = ref({});
    const df =
      getCurrentInstance().appContext.config.globalProperties.$df.value;
    const selectOptions = [
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

    const updateIfSelect = (value) => {
      updateNode(dataNode, nodeId, df, value, 'operator');
    };

    onMounted(async () => {
      await nextTick();
      nodeId.value = if_ref.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);
      operator.value = dataNode.value.data.operator;
    });

    return {
      if_ref,
      selectOptions,
      operator,
      updateIfSelect,
    };
  },
});
</script>
