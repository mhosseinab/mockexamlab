'use client'

import React from 'react'
import QuestionContainer from 'container/QuessionContainer'
import Instruction from './Instruction'
import { useSearchParams } from 'next/navigation'
import WritingQuestion from './WritingQuestion'
import { useTestsStore } from 'store/testsStore'

const Writing = () => {
  const test = useTestsStore((s) => s.test)
  const searchParams = useSearchParams()
  const level = searchParams.get('level')

  const writingTest = test?.sections?.filter(
    (item: any) => item.componentType == 'Writing'
  ) ?? []

  if (level == 'intro')
    return (
      <QuestionContainer title=''>
        <Instruction title='Writing' time='30' name='writing' section='writing' />
      </QuestionContainer>
    )

  return (
    <QuestionContainer title=''>
      <WritingQuestion question={writingTest[0]} />
    </QuestionContainer>
  )
}

export default Writing
