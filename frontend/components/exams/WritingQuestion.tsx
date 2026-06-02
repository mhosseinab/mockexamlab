'use client'

import React, { FC } from 'react'
import style from 'styles/exams/questions.module.scss'
import { useTestsStore } from 'store/testsStore'

interface Props {
  question: any
}

const WritingQuestion: FC<Props> = ({ question }) => {
  const [writingAnswer, setWritingAnswer] = React.useState('')
  const { answerTest } = useTestsStore()

  const submitAnswer = () => {
    answerTest(
      writingAnswer,
      '1c954620-2f9f-4200-9c92-3ebb1283178c',
      '36e0897b-41c4-4d07-8657-76c56c15bb9d'
    )
  }

  return (
    <div className={style.writing}>
      <div className={style.writing__question}>
        <div className='sections__header'>
          <h4>Writing</h4>
        </div>
        <div className={style.question}>
          <p dangerouslySetInnerHTML={{ __html: question.questionGroups[0].writingTopic.topic }} />
        </div>
      </div>
      <div className={style.writing__answer}>
        <div className={style.title}>
          <h4>{question.title}</h4>
          <div className={style.br} />
        </div>
        <h5>Write at least {question.sectionWriting.minimumWords} words.</h5>
        <textarea value={writingAnswer} onChange={(event) => setWritingAnswer(event.target.value)} rows={25} />
        <button onClick={submitAnswer}>Submit</button>
      </div>
    </div>
  )
}

export default WritingQuestion
