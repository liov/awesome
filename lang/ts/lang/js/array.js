
function isArray(obj) {
   return  Object.prototype.toString.call(obj).slice(8, -1) === 'Array'
}

function isArray2(obj) {
   return  Array.isArray(obj)
}

const arr = [1, 2, 3];
arr.length = 0;
console.log(arr); // 输出: []