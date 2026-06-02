'use client'

import React, { FC, useEffect } from 'react'
import style from 'styles/dashboard/exam.module.scss'
import { useTestsStore } from 'store/testsStore'
import { toast } from 'react-toastify'

interface Props {
  value: string
}

const ExamResults: FC<Props> = ({ value }) => {
  const { tests, getTest } = useTestsStore()
  const [results, setResults] = React.useState<
    { name: string; id: string; module: string; description: string }[]
  >([])

  const createTestHandler = (id: string) => {
    if (id.length < 1) {
      return toast.error('Please select a exam')
    }
    getTest(id)
  }

  useEffect(() => {
    if (value === '') return
    setResults(
      (tests || []).filter((item: any) => item.name.toLowerCase().includes(value))
    )
  }, [tests, value])

  if (!tests) return <h4>exams dose not exist.</h4>

  return (
    <div className={style.exams__results}>
      {value.length < 1
        ? (tests as any)?.data?.map(
            (item: { name: string; id: string; module: string; description: string }) => (
              <div onClick={() => createTestHandler(item.id)} className={style.exam} key={item.id}>
                <h4>{item.name.toUpperCase()}</h4>
                <span>level: {item.module}</span>
              </div>
            )
          )
        : results.map((item, index) => (
            <div className={style.exam} key={index}>
              <h4>{item.name.toUpperCase()}</h4>
              <span>level: {item.module}</span>
            </div>
          ))}
    </div>
  )
}

export default ExamResults
