<template>
  <div ref="el" class="add">
    <h4>Else</h4>
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
export default defineComponent({});
</script>
<style scoped>
.add {
  border-radius: 1rem;
  background-color: #0099de;
  padding: 1rem;
}
</style>
