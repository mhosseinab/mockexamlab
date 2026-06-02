'use client'

import ExamContainer from 'container/ExamContainer'
import Spinner from 'components/ui/Spinner'
import React, { useEffect } from 'react'
import { useAuthStore } from 'store/authStore'
import { useTestsStore } from 'store/testsStore'

export default function ExamsPage() {
  const [loading, setLoading] = React.useState(false)
  const { token } = useAuthStore()
  const { getAllExams } = useTestsStore()

  useEffect(() => {
    setLoading(true)
    if (token) getAllExams(token)
    setTimeout(() => { setLoading(false) }, 5000)
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  return !loading ? <ExamContainer /> : <Spinner />
}
