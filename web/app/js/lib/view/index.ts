/**
 * @param {Node} node
 *
 * @returns {Node}
 */
export function removeChildNodes(node) {
  while (node.firstChild) {
    if (node.firstChild.childNodes.length > 0) {
      removeChildNodes(node.firstChild);
    }
    node.removeChild(node.firstChild);
  }
  return node;
}

export function findChild(parent, child) {}

/**
 * @param {Node}                  root
 * @param {Node|DocumentFragment} node
 *
 * @returns {Node}
 */
export function attachFragment(root, node) {
  removeChildNodes(root);

  const nodes = node instanceof Node ? [node] : node;

  nodes.forEach(node => root.appendChild(node));

  return root;
}

/**
 * @param {Node}                  root
 * @param {Node|DocumentFragment} node
 *
 * @returns {Node}
 */
export function appendFragment(root, node) {
  const nodes = node instanceof Node ? [node] : node;

  nodes.forEach(node => root.appendChild(node));

  return root;
}
