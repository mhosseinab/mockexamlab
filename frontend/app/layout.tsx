import type { Metadata } from 'next'
import '../styles/style.scss'
import Providers from 'components/Providers'

export const metadata: Metadata = {
  title: 'Mockxlab',
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>
        <Providers>{children}</Providers>
      </body>
    </html>
  )
}
