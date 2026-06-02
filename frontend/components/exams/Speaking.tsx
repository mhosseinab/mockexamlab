'use client'

import React, { FC } from 'react'
import QuestionContainer from 'container/QuessionContainer'
import { useSearchParams } from 'next/navigation'
import Instruction from './Instruction'
import { useTestsStore } from 'store/testsStore'
import SpeakingQuestion from './SpeakingQuestion'

const Speaking: FC = () => {
  const test = useTestsStore((s) => s.test)
  const searchParams = useSearchParams()
  const level = searchParams.get('level')

  const speakingTest = test?.sections?.filter(
    (speak: any) => speak.componentType === 'Speaking'
  ) ?? []

  if (level == 'intro') {
    return (
      <QuestionContainer title=''>
        <Instruction title={speakingTest[0]?.componentType} time='30' name='' section='speaking' />
      </QuestionContainer>
    )
  }

  return (
    <QuestionContainer title={speakingTest[0]?.componentType ?? ''}>
      {speakingTest?.map((question: any, index: number) => (
        <SpeakingQuestion
          key={index}
          title={question.title}
          question={question.questionGroups[0]?.SpeakingTopic.topic}
          qNumber={question.questionGroups[0]?.SpeakingTopic.q_number}
        />
      ))}
    </QuestionContainer>
  )
}

export default Speaking
