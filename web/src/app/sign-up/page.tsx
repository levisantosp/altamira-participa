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

const schema = z
  .object({
    displayName: z
      .string()
      .min(1, 'Informe o nome completo')
      .max(100, 'O nome deve ter no máximo 100 caracteres'),
    username: z
      .string()
      .min(3, 'O nome de usuário deve ter no mínimo 3 caracteres')
      .max(32, 'O nome de usuário deve ter no máximo 32 caracteres')
      .regex(
        /^[a-zA-Z0-9_]+$/,
        'O nome de usuário só pode conter letras, números e underscore'
      ),
    email: z.email('Informe um e-mail válido'),
    password: z
      .string()
      .min(8, 'A senha deve ter no mínimo 8 caracteres')
      .max(72, 'A senha deve ter no máximo 72 caracteres'),
    confirmPassword: z.string().min(1, 'Confirme a senha')
  })
  .refine((data) => data.password === data.confirmPassword, {
    error: 'As senhas não coincidem',
    path: ['confirmPassword']
  })

export default function SignUp() {
  const form = useForm<z.infer<typeof schema>>({
    resolver: zodResolver(schema)
  })

  const router = useRouter()

  async function signUp(data: z.infer<typeof schema>) {
    try {
      await auth.signUp.email({
        email: data.email,
        password: data.password,
        username: data.username,
        displayName: data.displayName
      })

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
    <div className='flex min-h-[calc(100vh-8.75rem)] items-center justify-center w-full'>
      <Card className='w-full max-w-md md:max-w-xl'>
        <CardHeader>
          <CardTitle>Crie sua conta</CardTitle>
          <CardDescription>
            Preencha os dados abaixo para se cadastrar
          </CardDescription>

          <CardAction>
            <Link href='/sign-in'>
              <Button variant='link'>Entrar</Button>
            </Link>
          </CardAction>
        </CardHeader>

        <CardContent>
          <form id='sign-up' onSubmit={form.handleSubmit(signUp)}>
            <div className='flex flex-col gap-6'>
              <div className='grid gap-2'>
                <Label htmlFor='displayName'>Nome completo</Label>
                <Input
                  id='displayName'
                  type='text'
                  placeholder='Seu nome'
                  {...form.register('displayName')}
                />

                {form.formState.errors.displayName?.message && (
                  <span className='text-sm text-red-400'>
                    {form.formState.errors.displayName.message}
                  </span>
                )}
              </div>
              <div className='grid gap-2'>
                <Label htmlFor='username'>Nome de usuário</Label>
                <Input
                  id='username'
                  type='text'
                  {...form.register('username')}
                />

                {form.formState.errors.username?.message && (
                  <span className='text-sm text-red-400'>
                    {form.formState.errors.username.message}
                  </span>
                )}
              </div>
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
                <Label htmlFor='password'>Senha</Label>
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
              <div className='grid gap-2'>
                <Label htmlFor='confirmPassword'>Confirmar senha</Label>
                <Input
                  id='confirmPassword'
                  type='password'
                  {...form.register('confirmPassword')}
                />

                {form.formState.errors.confirmPassword?.message && (
                  <span className='text-sm text-red-400'>
                    {form.formState.errors.confirmPassword.message}
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
            form='sign-up'
          >
            {form.formState.isSubmitting ? (
              <Spinner />
            ) : (
              <span>Criar conta</span>
            )}
          </Button>
        </CardFooter>
      </Card>
    </div>
  )
}
