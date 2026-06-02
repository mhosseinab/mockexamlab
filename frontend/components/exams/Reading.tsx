'use client'

import React from 'react'
import Instruction from './Instruction'
import QuestionContainer from 'container/QuessionContainer'
import { useAuthStore } from 'store/authStore'
import { useTestsStore } from 'store/testsStore'
import ReadingQuestion from './ReadingQuesstion'
import { useSearchParams } from 'next/navigation'

const Reading = () => {
  const userName = useAuthStore((s) => s.userName)
  const test = useTestsStore((s) => s.test)
  const searchParams = useSearchParams()
  const level = searchParams.get('level')

  return (
    <QuestionContainer title={userName ?? ''}>
      {level == 'intro' ? (
        <Instruction title='Instruction' time='30' name='IELTS Reading' section='reading' />
      ) : (
        <div>
          {test?.sections && <ReadingQuestion />}
        </div>
      )}
    </QuestionContainer>
  )
}

export default Reading
