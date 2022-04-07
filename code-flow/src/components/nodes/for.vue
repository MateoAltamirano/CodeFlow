<template>
  <div ref="for_ref">
    <node title="For" color="#ae59b9">
      <el-input v-model="variable" @change="updateForInput" size="small" />
      <el-input-number
        v-model.number="start"
        @change="updateStartInput"
        controls-position="right"
        size="small"
      />
      <el-input-number
        v-model.number="stop"
        @change="updateStopInput"
        controls-position="right"
        size="small"
      />
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
    const for_ref = ref(null);
    const nodeId = ref(0);
    const variable = ref(null);
    const start = ref(null);
    const stop = ref(null);
    const dataNode = ref({});
    const df =
      getCurrentInstance().appContext.config.globalProperties.$df.value;

    const updateForInput = (value) => {
      updateNode(dataNode, nodeId, df, value, 'variable');
    };

    const updateStartInput = (value) => {
      updateNode(dataNode, nodeId, df, value, 'start');
    };

    const updateStopInput = (value) => {
      updateNode(dataNode, nodeId, df, value, 'stop');
    };

    onMounted(async () => {
      await nextTick();
      nodeId.value = for_ref.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);

      variable.value = dataNode.value.data.variable;
      start.value = dataNode.value.data.start;
      stop.value = dataNode.value.data.stop;
    });

    return {
      for_ref,
      variable,
      start,
      stop,
      updateForInput,
      updateStartInput,
      updateStopInput,
    };
  },
});
</script>
