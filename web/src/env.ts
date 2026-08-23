import { z } from 'zod/mini'

const schema = z.object({
  NEXT_PUBLIC_API_URL: z.url()
})

export const env = schema.parse(process.env)
