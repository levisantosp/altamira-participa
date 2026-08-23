import { z } from 'zod/mini'

const schema = z.object({
  NEXT_PUBLIC_API_URL: z.url(),
  NEXT_PUBLIC_ENABLE_API_DELAY: z.stringbool()
})

export const env = schema.parse(process.env)
