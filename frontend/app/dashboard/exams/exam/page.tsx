'use client'

import { Suspense, useEffect } from 'react'
import { useSearchParams } from 'next/navigation'
import { getCookie } from 'cookies-next'
import { useTestsStore } from 'store/testsStore'
import Spinner from 'components/ui/Spinner'
import Reading from 'components/exams/Reading'
import Listening from 'components/exams/Listening'
import Speaking from 'components/exams/Speaking'
import Writing from 'components/exams/Writing'

function ExamContent() {
  const searchParams = useSearchParams()
  const { createdTest, test, getTest } = useTestsStore()
  const section = searchParams.get('section')

  useEffect(() => {
    if (!test.id) {
      getTest(createdTest.testId || (getCookie('testId') as string))
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [test.id, createdTest.testId])

  if (!test.id) return <Spinner />

  switch (section?.toLowerCase()) {
    case 'reading': return <Reading />
    case 'listening': return <Listening />
    case 'speaking': return <Speaking />
    case 'writing': return <Writing />
    default: return <Spinner />
  }
}

export default function ExamPage() {
  return (
    <Suspense fallback={<Spinner />}>
      <ExamContent />
    </Suspense>
  )
}
