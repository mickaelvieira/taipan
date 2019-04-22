/**
* @param {Array} keys
* @param {Array} values

* @returns {Object}
*/
export default function(keys: string[], values: []) {
  return keys.reduce(function(carry, item, index) {
    return { ...carry, [item]: values[index] };
  }, {});
}
