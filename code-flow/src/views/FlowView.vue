<template>
  <div class="flow">
    <el-button type="primary" @click="exportEditor">Export</el-button>
    <el-aside width="250px" class="column">
      <div
        v-for="n in listNodes"
        :key="n"
        draggable="true"
        :data-node="n.item"
        @dragstart="drag($event)"
        class="drag-drawflow"
      >
        <div class="node" :style="`background: ${n.color}`">
          {{ n.name }}
        </div>
      </div>
    </el-aside>
    <el-main class="editor">
      <div
        id="drawflow"
        @drop="drop($event)"
        @dragover="allowDrop($event)"
      ></div>
    </el-main>
  </div>
  <el-dialog v-model="dialogVisible" title="Export" width="50%">
    <span>Data:</span>
    <pre><code>{{dialogData}}</code></pre>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="dialogVisible = false"
          >Confirm</el-button
        >
      </span>
    </template>
  </el-dialog>
</template>
<script>
import Drawflow from 'drawflow';
import styleDrawflow from 'drawflow/dist/drawflow.min.css';
import style from '../assets/style.css';
import {
  onMounted,
  shallowRef,
  h,
  getCurrentInstance,
  render,
  readonly,
  ref,
} from 'vue';
import Node1 from '../components/nodes/node1.vue';
import Node2 from '../components/nodes/node2.vue';
import Node3 from '../components/nodes/node3.vue';
import Number from '../components/nodes/number.vue';
import Add from '../components/nodes/add.vue';
import Assign from '../components/nodes/assign.vue';
import AST from '../models/ast.js';

export default {
  name: 'drawflow',
  setup() {
    const listNodes = readonly([
      {
        name: 'Number',
        color: '#00cd6a',
        item: 'Number',
        input: 0,
        output: 1,
        defaultData: { value: 0 },
      },
      {
        name: 'Add',
        color: '#ffc500',
        item: 'Add',
        input: 2,
        output: 1,
        defaultData: { operator: '+' },
      },
      {
        name: 'Assign',
        color: '#f15300',
        item: 'Assign',
        input: 1,
        output: 0,
        defaultData: { operator: '=', variable: 'x' },
      },
    ]);

    const editor = shallowRef({});
    const dialogVisible = ref(false);
    const dialogData = ref({});
    const Vue = { version: 3, h, render };
    const internalInstance = getCurrentInstance();
    internalInstance.appContext.app._context.config.globalProperties.$df =
      editor;

    function exportEditor() {
      dialogData.value = editor.value.export();
      const data = editor.value.export().drawflow.Home.data;
      parseData(data);
      dialogVisible.value = true;
    }
    function parseData(data) {
      let startingNodes;
      startingNodes = Object.values(data).filter(
        (node) => node.class === 'Assign'
      );
      if (startingNodes.length === 0) {
        startingNodes = Object.values(data).filter(
          (node) => node.class === 'Add'
        );
        if (startingNodes.length === 0) {
          startingNodes = Object.values(data).filter(
            (node) => node.class === 'Number'
          );
        }
      }

      const expressions = [];

      startingNodes.forEach((node) => {
        const ast = new AST(node, data);
        let expression = [];
        inOrderTraversal(ast, expression);
        expressions.push(expression.join(''));
        console.log(expression.join(''));
      });

      console.log({ expressions });

      function inOrderTraversal(ast, expression) {
        if (ast !== null) {
          inOrderTraversal(ast.left, expression);
          if (ast.value.class === undefined)
            expression.push(`${ast.value.data.variable} `);
          else if (ast.value.class === 'Number')
            expression.push(`${ast.value.data.value} `);
          else expression.push(`${ast.value.data.operator} `);
          inOrderTraversal(ast.right, expression);
        }
      }
    }

    const drag = (ev) => {
      if (ev.type === 'touchstart') {
        mobile_item_selec = ev.target
          .closest('.drag-drawflow')
          .getAttribute('data-node');
      } else {
        ev.dataTransfer.setData('node', ev.target.getAttribute('data-node'));
      }
    };
    const drop = (ev) => {
      if (ev.type === 'touchend') {
        var parentdrawflow = document
          .elementFromPoint(
            mobile_last_move.touches[0].clientX,
            mobile_last_move.touches[0].clientY
          )
          .closest('#drawflow');
        if (parentdrawflow != null) {
          addNodeToDrawFlow(
            mobile_item_selec,
            mobile_last_move.touches[0].clientX,
            mobile_last_move.touches[0].clientY
          );
        }
        mobile_item_selec = '';
      } else {
        ev.preventDefault();
        var data = ev.dataTransfer.getData('node');
        addNodeToDrawFlow(data, ev.clientX, ev.clientY);
      }
    };
    const allowDrop = (ev) => {
      ev.preventDefault();
    };
    let mobile_item_selec = '';
    let mobile_last_move = null;
    function positionMobile(ev) {
      mobile_last_move = ev;
    }
    function addNodeToDrawFlow(name, pos_x, pos_y) {
      pos_x =
        pos_x *
          (editor.value.precanvas.clientWidth /
            (editor.value.precanvas.clientWidth * editor.value.zoom)) -
        editor.value.precanvas.getBoundingClientRect().x *
          (editor.value.precanvas.clientWidth /
            (editor.value.precanvas.clientWidth * editor.value.zoom));
      pos_y =
        pos_y *
          (editor.value.precanvas.clientHeight /
            (editor.value.precanvas.clientHeight * editor.value.zoom)) -
        editor.value.precanvas.getBoundingClientRect().y *
          (editor.value.precanvas.clientHeight /
            (editor.value.precanvas.clientHeight * editor.value.zoom));

      const nodeSelected = listNodes.find((ele) => ele.item == name);
      editor.value.addNode(
        name,
        nodeSelected.input,
        nodeSelected.output,
        pos_x,
        pos_y,
        name,
        nodeSelected.defaultData,
        name,
        'vue'
      );
    }
    onMounted(() => {
      var elements = document.getElementsByClassName('drag-drawflow');
      for (var i = 0; i < elements.length; i++) {
        elements[i].addEventListener('touchend', drop, false);
        elements[i].addEventListener('touchmove', positionMobile, false);
        elements[i].addEventListener('touchstart', drag, false);
      }

      const id = document.getElementById('drawflow');
      editor.value = new Drawflow(
        id,
        Vue,
        internalInstance.appContext.app._context
      );
      editor.value.start();

      editor.value.registerNode('Node1', Node1, {}, {});
      editor.value.registerNode('Node2', Node2, {}, {});
      editor.value.registerNode('Node3', Node3, {}, {});
      editor.value.registerNode('Number', Number, { value: 0 }, {});
      editor.value.registerNode('Add', Add, {}, {});
      editor.value.registerNode('Assign', Assign, {}, {});
      const df = internalInstance.appContext.config.globalProperties.$df.value;
      editor.value.on(
        'connectionCreated',
        ({ output_id, input_id, output_class, input_class }) => {
          const inputNode = df.getNodeFromId(output_id);
          const outputNode = df.getNodeFromId(input_id);
          if (outputNode.class === 'Add') {
          }
        }
      );
      editor.value.import({
        drawflow: {
          Home: {
            data: {},
          },
        },
      });
    });
    return {
      exportEditor,
      listNodes,
      drag,
      drop,
      allowDrop,
      dialogVisible,
      dialogData,
    };
  },
};
</script>
<style scoped>
.flow {
  display: flex;
  height: 100%;
}

.column {
}

.editor {
  padding: 0;
}

.node {
  border-radius: 8px;
  border: 2px solid #494949;
  display: block;
  height: 60px;
  line-height: 40px;
  padding: 10px;
  cursor: move;
}

#drawflow {
  width: 100%;
  height: 100%;
  background: #1e3c50;
}
</style>
