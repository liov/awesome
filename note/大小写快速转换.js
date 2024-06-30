// 这里仅做演示,js直接toUpperCase()更简单易读
// 其他语言支持byte(uint8)的语言,可以直接'z' ^ ' '

const z = 'z'.charCodeAt(0);
console.log(String.fromCharCode(z ^ ' '.charCodeAt(0)));