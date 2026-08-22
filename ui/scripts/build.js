import { readdirSync, writeFileSync } from 'node:fs'

let content = ''

for (const file of readdirSync('./src/components')) {
  content += `export * from './components/${file.replace('.tsx', '')}'\n`
}

writeFileSync('./src/index.ts', content, 'utf-8')
