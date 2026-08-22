import { pluginAxios } from '@kubb/plugin-axios'
import { pluginReactQuery } from '@kubb/plugin-react-query'
import { pluginTs } from '@kubb/plugin-ts'
import { pluginZod } from '@kubb/plugin-zod'
import { defineConfig } from 'kubb/config'

export default defineConfig({
  input: 'http://localhost:3333/openapi.json',
  output: {
    path: './src/gen',
    clean: true
  },
  plugins: [
    pluginTs(),
    pluginAxios({
      baseURL: 'http://localhost:3333'
    }),
    pluginReactQuery({
      client: 'axios'
    }),
    pluginZod()
  ]
})
