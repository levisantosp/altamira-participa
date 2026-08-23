'use client'

import { zodResolver } from '@hookform/resolvers/zod'
import { ResponseError } from 'api-client'
import Link from 'next/link'
import { useRouter } from 'next/navigation'
import { useForm } from 'react-hook-form'
import { toast } from 'sonner'
import {
  Button,
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
  Input,
  Label,
  Spinner
} from 'ui'
import { z } from 'zod'
import { auth } from '@/lib/auth'

const schema = z.object({
  email: z.email('Informe um e-mail válido'),
  password: z.string().min(1, 'Informe a senha')
})

export default function SignIn() {
  const form = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema)
  })

  const router = useRouter()

  async function signIn(data: z.infer<typeof schema>) {
    try {
      await auth.signIn.email(data.email, data.password)
      router.push('/')
    } catch (err) {
      console.error(err)
      if (err instanceof ResponseError) {
        toast.error(err.response.data.detail)
      } else if (err instanceof Error) {
        toast.error('Ocorreu um erro inesperado', {
          description: err.message
        })
      }
    }
  }

  return (
    <>
      <Card className='w-full max-w-xl'>
        <CardHeader>
          <CardTitle>Bem vindo(a) de volta!</CardTitle>
          <CardDescription>
            Informe seu e-mail e senha abaixo para entrar na sua conta
          </CardDescription>

          <CardAction>
            <Link href='/sign-up'>
              <Button variant='link'>Criar conta</Button>
            </Link>
          </CardAction>
        </CardHeader>

        <CardContent>
          <form id='sign-in' onSubmit={form.handleSubmit(signIn)}>
            <div className='flex flex-col gap-6'>
              <div className='grid gap-2'>
                <Label htmlFor='email'>E-mail</Label>
                <Input
                  id='email'
                  type='email'
                  placeholder='usuario@exemplo.com'
                  {...form.register('email')}
                />

                {form.formState.errors.email?.message && (
                  <span className='text-sm text-red-400'>
                    {form.formState.errors.email.message}
                  </span>
                )}
              </div>
              <div className='grid gap-2'>
                <div className='flex items-center'>
                  <Label htmlFor='password'>Password</Label>
                  <a
                    href='#'
                    className='ml-auto inline-block text-sm underline-offset-4 hover:underline'
                  >
                    Esqueci minha senha
                  </a>
                </div>
                <Input
                  id='password'
                  type='password'
                  {...form.register('password')}
                />

                {form.formState.errors.password?.message && (
                  <span className='text-sm text-red-400'>
                    {form.formState.errors.password.message}
                  </span>
                )}
              </div>
            </div>
          </form>
        </CardContent>

        <CardFooter className='flex-col gap-2'>
          <Button
            type='submit'
            className='w-full'
            disabled={form.formState.isSubmitting}
            form='sign-in'
          >
            {form.formState.isSubmitting ? <Spinner /> : <span>Entrar</span>}
          </Button>
          <Button variant='outline' className='w-full' disabled>
            Login with Google
          </Button>
        </CardFooter>
      </Card>
    </>
  )
}
