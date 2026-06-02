'use client'

import Image from 'next/image'
import Link from 'next/link'
import React, { FC, useEffect } from 'react'
import style from 'styles/menu.module.scss'
import { getCookie } from 'cookies-next'

interface Props {
  links: {
    title: string
    link: string
  }[]
}

const Navbar: FC<Props> = ({ links }) => {
  const allLinks = getCookie('token') ? [{ link: '/dashboard', title: 'Dashboard' }] : links
  const [isSSR, setIsSSR] = React.useState(true)

  useEffect(() => {
    setIsSSR(false)
  }, [])

  return (
    <nav className={style.navbar}>
      <div className={style.logo}>
        <Image src='/img/logo.png' alt='logo' fill />
      </div>
      <ul className={style.menu}>
        <li>
          <Link href='/exams'>
            <Image src='/img/icons/save.png' alt='icon' width={16} height={18} />
            Exams
          </Link>
        </li>
        <li>
          <Link href='/courses'>
            <Image src='/img/icons/video.png' alt='icon' width={16} height={18} />
            Courses
          </Link>
        </li>
        <li>
          <Link href='/shop'>
            <Image src='/img/icons/shop.png' alt='icon' width={16} height={18} />
            Shop
          </Link>
        </li>
      </ul>
      {!isSSR && (
        <div className={style.auth__btn}>
          {allLinks.map((link, index) => (
            <Link key={index} href={link.link} className={style.btn}>{link.title}</Link>
          ))}
        </div>
      )}
    </nav>
  )
}

export default Navbar
