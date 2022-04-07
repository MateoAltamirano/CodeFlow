<template>
  <div class="node" :style="node_style">
    <h4>{{ title }}</h4>
    <slot></slot>
  </div>
</template>

<script>
import { defineComponent } from 'vue';
export default defineComponent({
  props: ['title', 'color'],

  setup(props) {
    const node_style = { backgroundColor: props.color };
    return { node_style };
  },
});
</script>

<style scoped>
.node {
  border-radius: 1rem;
  padding: 1rem;
}
</style>
