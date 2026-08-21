import { existsSync, readdirSync, unlinkSync, writeFileSync } from 'node:fs'

let content = ''

for (const file of readdirSync('./src/components')) {
  content += `export * from './components/${file.replace('.tsx', '')}'\n`
}

if (existsSync('./src/index.ts')) {
  unlinkSync('./src/index.ts')
}

writeFileSync('./src/index.ts', content, 'utf-8')
