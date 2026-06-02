'use client'

import React, { FC, Suspense } from 'react'
import { useSearchParams } from 'next/navigation'
import style from 'styles/exams/questions.module.scss'
import ExamsFooter from 'components/exams/Footer'
import TestHeader from 'components/exams/TestHeader'
import { useAuthStore } from 'store/authStore'
import { useTestsStore } from 'store/testsStore'

interface Props {
  children: React.ReactNode
  title: string
  playAudioHandler?: () => void
}

const QuestionContainerInner: FC<Props> = ({ children }) => {
  const userName = useAuthStore((s) => s.userName)
  const test = useTestsStore((s) => s.test)
  const searchParams = useSearchParams()
  const level = searchParams.get('level')

  return (
    <div className={style.question__layout}>
      <TestHeader title={userName ?? ''} />
      <div className={style.children}>{children}</div>
      {level !== 'intro' && test?.sections && <ExamsFooter sections={test.sections} />}
    </div>
  )
}

const QuestionContainer: FC<Props> = (props) => (
  <Suspense>
    <QuestionContainerInner {...props} />
  </Suspense>
)

export default QuestionContainer
