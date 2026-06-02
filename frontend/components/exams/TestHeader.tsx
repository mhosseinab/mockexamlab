'use client'

import React, { FC } from 'react'
import style from 'styles/exams/questions.module.scss'
import Image from 'next/image'
import Timer from './Timer'
import { useSearchParams } from 'next/navigation'
import AudioPlayer from './AudioPlayer'

interface Props {
  title: string
}

const TestHeader: FC<Props> = ({ title }) => {
  const [isPlaying, setISplaying] = React.useState(true)
  const THREE_DAYS_IN_MS = 1000 * 60 * 30
  const NOW_IN_MS = new Date().getTime()
  const searchParams = useSearchParams()
  const section = searchParams.get('section')
  const level = searchParams.get('level')

  return (
    <div className={style.question__layout}>
      <div className={style.header}>
        {section == 'listening' && level != 'intro' && (
          <AudioPlayer isPlaying={isPlaying} />
        )}
        <div className={style.test_title}>
          <div className={style.icon}>
            <Image src='/img/icons/person.svg' alt='person icon' width={14} height={20} />
          </div>
          <h3>{title}</h3>
        </div>
        <div className={style.exam__timer}>
          <div className={style.bar} />
          <div className={style.exam__time}>
            {level == 'intro' ? (
              <>30:00</>
            ) : (
              <Timer deadline={new Date(THREE_DAYS_IN_MS + NOW_IN_MS)} />
            )}
            <svg width='14' height='14' viewBox='0 0 14 14' fill='none' xmlns='http://www.w3.org/2000/svg'>
              <path fillRule='evenodd' clipRule='evenodd' d='M7.00033 13.8337C10.7743 13.8337 13.8337 10.7743 13.8337 7.00033C13.8337 3.22638 10.7743 0.166992 7.00033 0.166992C3.22638 0.166992 0.166992 3.22638 0.166992 7.00033C0.166992 10.7743 3.22638 13.8337 7.00033 13.8337ZM12.3337 7.00033C12.3337 9.94584 9.94584 12.3337 7.00033 12.3337C4.05481 12.3337 1.66699 9.94584 1.66699 7.00033C1.66699 4.05481 4.05481 1.66699 7.00033 1.66699C9.94584 1.66699 12.3337 4.05481 12.3337 7.00033ZM7.75033 3.66699C7.75033 3.25278 7.41454 2.91699 7.00033 2.91699C6.58611 2.91699 6.25033 3.25278 6.25033 3.66699V6.5861C6.25033 7.05022 6.4347 7.49534 6.76288 7.82353L7.80332 8.86399C8.09622 9.15688 8.57109 9.15688 8.86399 8.86399C9.15688 8.5711 9.15688 8.09623 8.86399 7.80333L7.82355 6.76287C7.77666 6.71599 7.75033 6.6524 7.75033 6.5861V3.66699Z' fill='#52C3FF' />
            </svg>
          </div>
        </div>
        <div className={style.exam__actions}>
          <div onClick={() => setISplaying(!isPlaying)} className={style.icon}>
            <svg width='18' height='18' viewBox='0 0 18 18' fill='none' xmlns='http://www.w3.org/2000/svg'>
              <path fillRule='evenodd' clipRule='evenodd' d='M2.025 15.975V2.025H5.175V15.975H2.025ZM0 1.35C0 0.604414 0.604416 0 1.35 0H5.85C6.59558 0 7.2 0.604416 7.2 1.35V16.65C7.2 17.3956 6.59558 18 5.85 18H1.35C0.604417 18 0 17.3956 0 16.65V1.35ZM12.825 15.975V2.025H15.975V15.975H12.825ZM10.8 1.35C10.8 0.604414 11.4044 0 12.15 0H16.65C17.3956 0 18 0.604416 18 1.35V16.65C18 17.3956 17.3956 18 16.65 18H12.15C11.4044 18 10.8 17.3956 10.8 16.65V1.35Z' fill='white' />
            </svg>
          </div>
        </div>
      </div>
    </div>
  )
}

export default TestHeader
