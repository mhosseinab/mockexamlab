import { create } from 'zustand'
import { persist } from 'zustand/middleware'
import axios from 'axios'
import { deleteCookie, setCookie } from 'cookies-next'
import { toast } from 'react-toastify'
import auth from 'lib/firebase'
import { createUserWithEmailAndPassword, signInWithEmailAndPassword } from 'firebase/auth'

const API = 'https://mel-api.go7.ir/api/v1/'

const authorization = async (accessToken: string) => {
  const { data } = await axios.post(
    `${API}user/auth`,
    {},
    { headers: { Authorization: accessToken } }
  )
  return data
}

interface AuthState {
  email: string | null
  token: string | null
  userName: string | null
  userId: string | null
  isLoading: boolean
  login: (email: string, password: string) => Promise<void>
  signUp: (name: string, family: string, email: string, password: string) => Promise<void>
  reset: () => void
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      email: null,
      token: null,
      userName: null,
      userId: null,
      isLoading: false,

      login: async (email, password) => {
        set({ isLoading: true })
        deleteCookie('token')
        try {
          const credential = await signInWithEmailAndPassword(auth, email, password)
          const accessToken = await credential.user.getIdToken()
          const data = await authorization(accessToken)
          if (data.userId) {
            set({
              email: credential.user.email,
              userName: credential.user.displayName,
              token: accessToken,
              userId: data.userId,
              isLoading: false,
            })
            setCookie('token', accessToken, { maxAge: 3600 })
            setCookie('userId', data.userId, { maxAge: 60 * 60 * 24 })
            setCookie('refreshToken', credential.user.refreshToken, { maxAge: 60 * 60 * 24 * 30 })
          }
        } catch (error: any) {
          toast.error(error.code?.replace(/\W/g, ' ') ?? 'Login failed')
          set({ isLoading: false })
        }
      },

      signUp: async (name, family, email, password) => {
        set({ isLoading: true })
        deleteCookie('token')
        try {
          const credential = await createUserWithEmailAndPassword(auth, email, password)
          const accessToken = await credential.user.getIdToken()
          const data = await authorization(accessToken)
          if (data.userId) {
            set({
              email: credential.user.email,
              userName: `${name} ${family}`,
              token: accessToken,
              userId: data.userId,
              isLoading: false,
            })
            setCookie('token', accessToken, { maxAge: 3600 })
            setCookie('userId', data.userId, { maxAge: 60 * 60 * 24 })
            setCookie('refreshToken', credential.user.refreshToken, { maxAge: 60 * 60 * 24 * 30 })
          }
        } catch (error: any) {
          toast.error(error.code?.replace(/\W/g, ' ') ?? 'Sign up failed')
          set({ isLoading: false })
        }
      },

      reset: () => set({ email: null, token: null, userName: null, userId: null, isLoading: false }),
    }),
    { name: 'auth-storage' }
  )
)
