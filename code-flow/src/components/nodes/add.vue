<template>
  <div ref="el" class="add">
    <h4>Add</h4>
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
import nodeHeader from './nodeHeader.vue';
export default defineComponent({
  components: {
    nodeHeader,
  },
  // setup() {
  //   const el = ref(null);
  //   const nodeId = ref(0);
  //   let df = null;
  //   const result = ref(null);
  //   const dataNode = ref({});

  //   df = getCurrentInstance().appContext.config.globalProperties.$df.value;

  //   onMounted(async () => {
  //     await nextTick();
  //     nodeId.value = el.value.parentElement.parentElement.id.slice(5);
  //     dataNode.value = df.getNodeFromId(nodeId.value);

  //     result.value = dataNode.value.data.result;
  //   });

  //   return {
  //     el,
  //     result,
  //     updateSelect,
  //   };
  // },
});
</script>
<style scoped>
.add {
  border-radius: 1rem;
  background-color: #ffc500;
  padding: 1rem;
}
</style>
