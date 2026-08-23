import { readdirSync, writeFileSync } from 'node:fs'

let content = `import { client } from './gen/.kubb/client'

client.interceptors.request.use((req) => {
  req.withCredentials = true
  return req
})

`

for (const folder of readdirSync('./src/gen')) {
  for (const file of readdirSync(`./src/gen/${folder}`)) {
    if (file === 'index.ts') continue
    content += `export * from './gen/${folder}/${file}'\n`
  }
}

writeFileSync('./src/index.ts', content, 'utf-8')
