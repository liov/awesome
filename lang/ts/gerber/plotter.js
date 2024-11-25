import fs from 'fs'
import path from 'path'
import { parse } from '@tracespace/parser'
import {plot} from '@tracespace/plotter'

// 读取配置文件
const configPath = path.join(__dirname, 'config.json');
const config = JSON.parse(fs.readFileSync(configPath, 'utf8'));

const gerberContents = await fs.readFile(String.raw`xxx`, 'utf-8')
const syntaxTree = parse(gerberContents)
const imageTree = plot(syntaxTree)

console.log(JSON.stringify(imageTree, null, 2))