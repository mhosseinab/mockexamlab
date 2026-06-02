import HomeLayout from 'components/ui/HomeLayout'
import Navbar from 'components/ui/Navbar'
import HomeLanding from 'container/HomeLanding'

export default function Home() {
  return (
    <>
      <HomeLayout>
        <Navbar links={[{ link: '/signup', title: 'Sign up' }, { link: '/signin', title: 'Sign in' }]} />
      </HomeLayout>
      <HomeLanding />
    </>
  )
}
