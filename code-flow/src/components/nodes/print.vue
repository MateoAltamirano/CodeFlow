<template>
  <div ref="el" class="add">
    <h4>Print</h4>
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
  // setup() {
  //   const el = ref(null);
  //   const nodeId = ref(0);
  //   let df = null;
  //   const dataNode = ref({});
  //   df = getCurrentInstance().appContext.config.globalProperties.$df.value;
  //   return {
  //     el,
  //   };
  // },
});
</script>
<style scoped>
.add {
  border-radius: 1rem;
  background-color: #8fa7a6;
  padding: 1rem;
}
</style>
