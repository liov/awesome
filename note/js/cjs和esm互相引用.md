# cjs引入esm
```js
// https://github.com/nodejs/node/pull/51977 --experimental-require-module 
const module = require('./esm.mjs')
```
你可以使用`import`语句来导入ES Module，而不是使用`require`。
```js
async function add_module() {
    const module = await import( './esm.mjs');
}
```
```js
import('./esm.mjs').then(module => {})
```
# esm引入cjs
```js
import module from './cjs.cjs';
```
```js
// Deprecated 
import {createRequire} from 'node:module'
const require = createRequire(import.meta.url);
const module = require('./cjs.cjs');
```

使用 .cjs 扩展名：
如果你的 CJS 模块文件使用 .cjs 扩展名，那么在 ESM 代码中可以直接使用 import 语句引入它。
使用 .js 扩展名和 import() 动态导入：
如果你的 CJS 模块使用 .js 扩展名，那么可以使用 import() 动态导入函数来引入它。
使用 .json 扩展名：
如果你的 CJS 模块是一个 JSON 文件，可以直接使用 import 语句来引入。

意事项
默认导出：如果 ESM 模块使用 export default 导出一个对象，那么在 CJS 中，你可以直接使用 require 语句来获取这个默认导出的对象。
命名导出：如果 ESM 模块使用 export 关键字导出多个命名成员，那么在 CJS 中，你需要使用 require 语句，并且通过访问 __esModule 属性来获取这些命名导出。这是因为 CJS 模块系统没有直接支持命名导出。
```js
// index.js
const example = require('./example.mjs');

if (example && example.__esModule) {
  console.log(example.greet('World'));  // 输出 "Hello, World!"
  console.log(example.farewell('World'));  // 输出 "Goodbye, World!"
} else {
  // 如果模块没有使用 __esModule，那么它可能是 CommonJS 模块
  console.log(example.greet('World'));  // 输出 "Hello, World!"
}
```