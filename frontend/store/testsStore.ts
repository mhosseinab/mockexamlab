import { create } from 'zustand'
import axios from 'axios'
import { getCookie, setCookie } from 'cookies-next'
import { toast } from 'react-toastify'

const API = 'https://mel-api.go7.ir/api/v1/'

interface Test {
  id: string
  name: string
  description: string
  testDate: string
  section: string
}

interface CreatedTest {
  id: string
  testId?: string
  name: string
  description: string
  module: string
  testDate: string
  createdAt: string
}

interface TestsState {
  tests: Test[] | null
  createdTest: CreatedTest
  test: any
  getAllExams: (token: string) => Promise<void>
  createTest: (testId: string, testType: string, testDate: Date, onSuccess?: () => void) => Promise<void>
  getTest: (testId: string) => Promise<void>
  answerTest: (answer: string, questionId: string, userTestId: string) => Promise<void>
}

export const useTestsStore = create<TestsState>()((set) => ({
  tests: null,
  createdTest: { id: '', name: '', description: '', module: '', testDate: '', createdAt: '' },
  test: {},

  getAllExams: async (token) => {
    try {
      const { data } = await axios.get(`${API}test/all`, {
        headers: { Authorization: token },
      })
      set({ tests: data })
    } catch {
      toast.error('Cannot get exams, try again')
    }
  },

  createTest: async (testId, _testType, _testDate, onSuccess) => {
    try {
      const { data } = await axios.post(
        `${API}user-test`,
        JSON.stringify({ testDate: 0, testId, testType: 0, userId: getCookie('userId') }),
        {
          headers: {
            Authorization: getCookie('token') as string,
            'content-type': 'application/json',
          },
        }
      )
      setCookie('testId', testId, { maxAge: 60 * 60 * 24 })
      toast.success('Your Exam successfully created')
      set({ createdTest: data })
      if (onSuccess) setTimeout(onSuccess, 3000)
    } catch (err: any) {
      toast.error(err.response?.data ?? 'Failed to create test')
    }
  },

  getTest: async (testId) => {
    try {
      const { data } = await axios.get(`${API}test/full/${testId}`, {
        headers: {
          Authorization: getCookie('token') as string,
          'content-type': 'application/json',
        },
      })
      set({ test: data })
    } catch (err: any) {
      toast.error(err.response?.data ?? 'Failed to get test')
    }
  },

  answerTest: async (answer, questionId, userTestId) => {
    try {
      await axios.post(
        `${API}user-answer`,
        JSON.stringify({ answer, questionId, userTestId }),
        {
          headers: {
            'content-type': 'application/json',
            Authorization: getCookie('token') as string,
          },
        }
      )
    } catch (err: any) {
      toast.error(err.response?.data ?? 'Failed to submit answer')
    }
  },
}))
