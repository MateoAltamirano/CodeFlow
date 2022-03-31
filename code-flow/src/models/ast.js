export default class AST {
  constructor(node, data, tabs) {
    this.value = node;
    if (node.class === 'If' || node.class === 'For') tabs += 1;
    this.left = this.getInputNode(node, data, 'input_1', tabs);
    this.right = this.getInputNode(node, data, 'input_2', tabs);
    this.tabs = tabs;
    if (node.class === 'Assign' || node.class === 'For') this.swapNodes();
  }

  getInputNode(node, data, inputNum, tabs) {
    if (node.class === 'Assign' && inputNum === 'input_2') {
      const { variable } = node.data;
      const newNodeVariable = { data: { variable }, inputs: {} };
      return new AST(newNodeVariable, data, tabs);
    }
    if (node.class === 'If') {
      if (inputNum === 'input_1') {
        const { operator } = node.data;
        const left =
          node.inputs.input_1.connections.length > 0
            ? data[node.inputs.input_1.connections[0].node].data.value
            : '';
        const right =
          node.inputs.input_2.connections.length > 0
            ? data[node.inputs.input_2.connections[0].node].data.value
            : '';
        const condition = `${left} ${operator} ${right}`;
        const newConditionNode = { data: { condition }, inputs: {} };
        return new AST(newConditionNode, data, tabs);
      } else {
        inputNum = 'input_3';
      }
    }
    if (node.class === 'For' && inputNum === 'input_2') {
      const { variable, start, stop } = node.data;
      const condition = `${variable} in range(${start}, ${stop})`;
      const newConditionNode = { data: { condition }, inputs: {} };
      return new AST(newConditionNode, data, tabs);
    }
    if (node.class === 'Print') {
      if (inputNum === 'input_1') {
        const childNode =
          node.inputs.input_1.connections.length > 0
            ? data[node.inputs.input_1.connections[0].node]
            : null;
        if (childNode && childNode.class) {
          if (
            childNode.class === 'Number' ||
            childNode.class === 'Variable' ||
            childNode.class === 'String' ||
            childNode.class === 'Boolean'
          ) {
            const value = childNode.data.value;
            const newValueNode = { data: { value }, inputs: {} };
            return new AST(newValueNode, data, tabs);
          }
        }
      }
      return null;
    }
    const nodeInput = node.inputs[inputNum];
    if (nodeInput === undefined || nodeInput['connections'].length === 0)
      return null;
    return new AST(data[nodeInput['connections'][0].node], data, tabs);
  }

  swapNodes() {
    const temp = this.left;
    this.left = this.right;
    this.right = temp;
  }
}
