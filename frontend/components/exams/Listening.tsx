'use client'

import React, { FC } from 'react'
import CheckQuestion from './CheckQuestion'
import { useTestsStore } from 'store/testsStore'
import { useAuthStore } from 'store/authStore'
import { useSearchParams } from 'next/navigation'
import Instruction from './Instruction'
import style from 'styles/exams/questions.module.scss'
import QuestionContainer from 'container/QuessionContainer'
import NoteCompletion from './NoteCompletion'

const Listening: FC = () => {
  const test = useTestsStore((s) => s.test)
  const userName = useAuthStore((s) => s.userName)
  const searchParams = useSearchParams()
  const level = searchParams.get('level')

  if (level == 'intro')
    return (
      <QuestionContainer title=''>
        <Instruction title='Instruction' time='30' name='IELTS Listening' section='listening' />
      </QuestionContainer>
    )

  const questions = test?.sections?.[1]?.questionGroups?.[1]?.multipleChoice || null
  const noteCompletionQuestions = test?.sections?.[1]?.questionGroups?.[0]

  return (
    <QuestionContainer title={userName ?? ''}>
      <div className={style.exam__title}>
        <h4>{test?.sections?.[2]?.componentType}</h4>
        <div className={style.br} />
      </div>
      <h4 className={style.exam__tests}>
        Questions {questions?.[0]?.q_number}- {questions?.[0]?.q_number + questions?.length}
      </h4>
      <p className={style.section__details}>{test?.sections?.[1]?.questionGroups?.[1]?.description}</p>
      {questions?.map((item: any, index: number) => (
        <CheckQuestion key={index} title={item.q_number + ' - ' + item.title} question={item.answers} />
      ))}
      {noteCompletionQuestions && (
        <>
          <h1>{noteCompletionQuestions.title}</h1>
          <p style={{ margin: '1em 0' }}>{noteCompletionQuestions.description}</p>
          <NoteCompletion question={noteCompletionQuestions.noteCompletion.note} />
        </>
      )}
    </QuestionContainer>
  )
}

export default Listening
