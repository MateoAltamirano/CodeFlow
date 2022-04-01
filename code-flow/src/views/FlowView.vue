<template>
  <div class="flow">
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
      <el-button
        color="#ebf0f1"
        style="color: gray"
        size="large"
        @click="generateCode"
        >Generate Code</el-button
      >
    </el-aside>
    <el-main class="editor">
      <div
        id="drawflow"
        @drop="drop($event)"
        @dragover="allowDrop($event)"
      ></div>
    </el-main>
  </div>
</template>
<script>
import Drawflow from 'drawflow';
import styleDrawflow from 'drawflow/dist/drawflow.min.css';
import style from '../assets/style.css';
import { useCodeFlowStore } from '../stores/codeFlow';
import { onBeforeRouteLeave } from 'vue-router';
import router from '../router';
import {
  onMounted,
  shallowRef,
  h,
  getCurrentInstance,
  render,
  readonly,
  ref,
} from 'vue';

import Variable from '../components/nodes/variable.vue';
import Number from '../components/nodes/number.vue';
import String from '../components/nodes/string.vue';
import Boolean from '../components/nodes/boolean.vue';
import Operation from '../components/nodes/operation.vue';
import Assign from '../components/nodes/assign.vue';
import If from '../components/nodes/if.vue';
import Else from '../components/nodes/else.vue';
import For from '../components/nodes/for.vue';
import Print from '../components/nodes/print.vue';
import AST from '../models/ast.js';

export default {
  name: 'drawflow',
  setup() {
    const codeFlowStore = useCodeFlowStore();
    const listNodes = readonly([
      {
        name: 'Variable',
        color: '#00a186',
        item: 'Variable',
        input: 0,
        output: 1,
        defaultData: { value: '' },
      },
      {
        name: 'Number',
        color: '#00cd6a',
        item: 'Number',
        input: 0,
        output: 1,
        defaultData: { value: 0 },
      },
      {
        name: 'String',
        color: '#00cd6a',
        item: 'String',
        input: 0,
        output: 1,
        defaultData: { value: '""' },
      },
      {
        name: 'Boolean',
        color: '#00cd6a',
        item: 'Boolean',
        input: 0,
        output: 1,
        defaultData: { value: 'True' },
      },
      {
        name: 'Addition',
        color: '#ffc500',
        item: 'Addition',
        input: 2,
        output: 1,
        defaultData: { operator: '+' },
      },
      {
        name: 'Substraction',
        color: '#ffc500',
        item: 'Substraction',
        input: 2,
        output: 1,
        defaultData: { operator: '-' },
      },
      {
        name: 'Multiplication',
        color: '#ffc500',
        item: 'Multiplication',
        input: 2,
        output: 1,
        defaultData: { operator: '*' },
      },
      {
        name: 'Division',
        color: '#ffc500',
        item: 'Division',
        input: 2,
        output: 1,
        defaultData: { operator: '/' },
      },
      {
        name: 'Assign',
        color: '#f15300',
        item: 'Assign',
        input: 1,
        output: 1,
        defaultData: { operator: '=', variable: 'x' },
      },
      {
        name: 'If',
        color: '#0099de',
        item: 'If',
        input: 3,
        output: 1,
        defaultData: { operator: '==' },
      },
      {
        name: 'Else',
        color: '#0099de',
        item: 'Else',
        input: 2,
        output: 1,
        defaultData: {},
      },
      {
        name: 'For',
        color: '#ae59b9',
        item: 'For',
        input: 1,
        output: 1,
        defaultData: { variable: 'i', start: 0, stop: 10 },
      },
      {
        name: 'Print',
        color: '#8fa7a6',
        item: 'Print',
        input: 1,
        output: 1,
        defaultData: {},
      },
    ]);

    const editor = shallowRef({});
    const Vue = { version: 3, h, render };
    const internalInstance = getCurrentInstance();
    internalInstance.appContext.app._context.config.globalProperties.$df =
      editor;

    function generateCode() {
      const flow = editor.value.export().drawflow.Home.data;
      const generatedCode = parseFlowToCode(flow);
      codeFlowStore.updateCode(generatedCode);
      router.push('/code');
    }

    function parseFlowToCode(data) {
      let rootNodes;
      rootNodes = Object.values(data).filter((node) => isRoot(node));

      const expressions = [];

      rootNodes.forEach((node) => {
        const ast = new AST(node, data, 0);
        let expression = [];
        inOrderTraversal(ast, expression);
        expressions.push(expression.join(''));
      });

      return expressions.join('\n');

      function inOrderTraversal(ast, expression) {
        if (ast !== null) {
          inOrderTraversal(ast.left, expression);
          switch (ast.value.class) {
            case 'Variable':
            case 'Number':
            case 'String':
            case 'Boolean':
              expression.push(`${ast.value.data.value} `);
              break;
            case 'Assign':
            case 'Addition':
            case 'Substraction':
            case 'Multiplication':
            case 'Division':
              expression.push(`${ast.value.data.operator} `);
              break;
            case 'If':
              expression.push(`if ${ast.left.value.data.condition}:\n`);
              for (let i = 0; i < ast.tabs; i++) {
                expression.push('\t');
              }
              break;
            case 'Else':
              expression.push('\n');
              for (let i = 0; i < ast.tabs - 1; i++) {
                expression.push('\t');
              }
              expression.push('else:\n');
              for (let i = 0; i < ast.tabs; i++) {
                expression.push('\t');
              }
              break;
            case 'For':
              expression.push(`for ${ast.left.value.data.condition}:\n`);
              for (let i = 0; i < ast.tabs; i++) {
                expression.push('\t');
              }
              break;
            case 'Print':
              const value = ast.left ? ast.left.value.data.value : '';
              expression.push(`print(${value})`);
              break;
            case undefined:
              if (
                ast.value.data.condition === undefined &&
                ast.value.data.value === undefined
              ) {
                expression.push(`${ast.value.data.variable} `);
              }
              break;
            default:
          }
          inOrderTraversal(ast.right, expression);
        }
      }

      function isRoot(node) {
        return node.outputs.output_1.connections.length === 0;
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

      editor.value.registerNode('Variable', Variable, {}, {});
      editor.value.registerNode('Number', Number, { value: 0 }, {});
      editor.value.registerNode('String', String, {}, {});
      editor.value.registerNode('Boolean', Boolean, {}, {});
      editor.value.registerNode('Addition', Operation, {}, {});
      editor.value.registerNode('Substraction', Operation, {}, {});
      editor.value.registerNode('Multiplication', Operation, {}, {});
      editor.value.registerNode('Division', Operation, {}, {});
      editor.value.registerNode('Assign', Assign, {}, {});
      editor.value.registerNode('If', If, {}, {});
      editor.value.registerNode('Else', Else, {}, {});
      editor.value.registerNode('For', For, {}, {});
      editor.value.registerNode('Print', Print, {}, {});

      editor.value.import({
        drawflow: {
          Home: {
            data: codeFlowStore.flow,
          },
        },
      });
    });

    onBeforeRouteLeave(() => {
      codeFlowStore.updateFlow(editor.value.export().drawflow.Home.data);
    });

    return {
      generateCode,
      listNodes,
      drag,
      drop,
      allowDrop,
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
  display: flex;
  flex-direction: column;
  overflow: auto;
  background-color: #23475f;
  padding: 1rem;
}

.editor {
  padding: 0;
  height: 100%;
}

.node {
  border-radius: 8px;
  border: 2px solid #494949;
  display: block;
  height: 60px;
  line-height: 40px;
  padding: 10px;
  cursor: move;
  margin-bottom: 1rem;
}

#drawflow {
  width: 100%;
  height: 100%;
  background: #1e3c50;
  background-size: 2rem 2rem;
  background-image: radial-gradient(circle, #798d8e 1px, rgba(0, 0, 0, 0) 1px);
}

.el-button {
  height: 60px;
  line-height: 40px;
  border-radius: 0.5rem;
  margin-bottom: 0;
}
</style>
