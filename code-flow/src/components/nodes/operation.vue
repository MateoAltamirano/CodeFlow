<template>
  <div ref="el" class="add">
    <h4>{{ name }}</h4>
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
    const name = ref(null);
    const dataNode = ref({});

    df = getCurrentInstance().appContext.config.globalProperties.$df.value;

    onMounted(async () => {
      await nextTick();
      nodeId.value = el.value.parentElement.parentElement.id.slice(5);
      dataNode.value = df.getNodeFromId(nodeId.value);

      name.value = dataNode.value.name;
    });

    return {
      el,
      name,
    };
  },
});
</script>
<style scoped>
.add {
  border-radius: 1rem;
  background-color: #ffc500;
  padding: 1rem;
}
</style>
