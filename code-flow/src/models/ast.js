export default class AST {
  constructor(node, data) {
    this.value = node;
    this.left = this.getInputNode(node, data, 'input_1');
    this.right = this.getInputNode(node, data, 'input_2');
    if (node.class === 'Assign') this.swapNodes();
  }

  getInputNode(node, data, inputNum) {
    if (node.class === 'Assign' && inputNum === 'input_2') {
      const { variable } = node.data;
      const newNodeVariable = { data: { variable }, inputs: {} };
      return new AST(newNodeVariable, data);
    }
    const nodeInput = node.inputs[inputNum];
    if (nodeInput === undefined || nodeInput['connections'].length === 0)
      return null;
    return new AST(data[nodeInput['connections'][0].node], data);
  }

  swapNodes() {
    const temp = this.left;
    this.left = this.right;
    this.right = temp;
  }
}
