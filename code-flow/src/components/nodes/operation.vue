<template>
  <div ref="operation_ref">
    <node :title="name" color="#ffc500"></node>
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

export default defineComponent({
  components: {
    node,
  },

  setup() {
    const operation_ref = ref(null);
    const nodeId = ref(0);
    const name = ref(null);
    const dataNode = ref({});
    const df =
      getCurrentInstance().appContext.config.globalProperties.$df.value;

    onMounted(async () => {
      await nextTick();
      nodeId.value =
        operation_ref.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);
      name.value = dataNode.value.name;
    });

    return {
      operation_ref,
      name,
    };
  },
});
</script>
