export const updateNode = (dataNode, nodeId, df, value, key) => {
  dataNode.value.data[key] = value;
  df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
};

export const onNodeMounted = (dataNode, nodeId, df, key) => {
  ref.value.parentElement.parentElement.id.slice(5);
  df.getNodeFromId(nodeId.value);
  dataNode.value.data[key];
};
