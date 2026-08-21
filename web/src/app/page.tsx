import { Button } from 'ui'

export default function Home() {
  return (
    <div className='min-h-screen flex justify-center items-center gap-5'>
      <Button variant='default'>Default</Button>
      <Button variant='secondary'>Secondary</Button>
      <Button variant='destructive'>Destructive</Button>
    </div>
  )
}
