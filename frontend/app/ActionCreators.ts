import { Actions } from "app/Actions";
import axios from "axios";
import { Dispatch } from "redux";
import { ActionTypes } from "./ActionTypes";
import Navigate from "next/router";
import { deleteCookie, getCookie, setCookie } from "cookies-next";
import { toast } from "react-toastify";
import auth from "lib/firebase";
import {
  createUserWithEmailAndPassword,
  signInWithEmailAndPassword,
  getIdToken,
} from "firebase/auth";

const api = "https://mel-api.go7.ir/api/v1/";

/*
################################################ Auth --get details Action Creator ################################################
*/

export const authorization = async (accessToken: string) => {
  const { data } = await axios.post(
    api + "user/auth",
    {},
    {
      headers: {
        Authorization: accessToken,
      },
    }
  );
  return data;
};

/*
################################################ Auth refresh token Action Creator ################################################
*/

const refreshToken = async () => {};

/*
################################################ Login Action Creator ################################################
*/
export const login = (email: string, password: string) => {
  return async (dispatch: Dispatch<Actions>) => {
    dispatch({
      type: ActionTypes.LOGIN,
      payload: {
        email: null,
        userName: null,
        userId: null,
        token: null,
        error: null,
        isLoading: true,
      },
    });
    deleteCookie("token");
    try {
      const response: any = await signInWithEmailAndPassword(
        auth,
        email,
        password
      );
      if (response._tokenResponse && response.user.accessToken) {
        const data = await authorization(response.user.accessToken); // get userId and alias and email from server
        if (data.userId) {
          const { displayName, expiresIn, idToken, email, refreshToken } =
            response._tokenResponse;
          dispatch({
            type: ActionTypes.LOGIN,
            payload: {
              email: email,
              userName: displayName,
              token: response.user.accessToken,
              userId: data.userId,
              error: null,
              isLoading: false,
            },
          });
          setCookie("token", response.user.accessToken, {
            maxAge: expiresIn,
          });
          setCookie("userId", data.userId, { maxAge: 60 * 60 * 24 });
          setCookie("refreshToken", refreshToken, {
            maxAge: 60 * 60 * 24 * 30,
          });
          await Navigate.push("/dashboard/exams");
        }
      }
    } catch (error: any) {
      toast.error(error.code.replace(/\W/g, " "));
      dispatch({
        type: ActionTypes.LOGIN,
        payload: {
          email: null,
          token: null,
          userName: null,
          userId: null,
          error: "",
          isLoading: false,
        },
      });
    }
  };
};

/*  ###########################################  Sign Up action Creator  #####################################################  */

export const signUp = (
  name: string,
  family: string,
  email: string,
  password: string
) => {
  return async (dispatch: Dispatch<Actions>) => {
    dispatch({
      type: ActionTypes.SIGN_UP,
      payload: {
        email: null,
        token: null,
        userName: null,
        error: null,
        userId: null,
        isLoading: true,
      },
    });
    deleteCookie("token");
    try {
      const response: any = await createUserWithEmailAndPassword(
        auth,
        email,
        password
      );
      if (response._tokenResponse && response.user.accessToken) {
        const data = await authorization(response.user.accessToken); // get userId and alias and email from server
        console.log(data);
        if (data.userId) {
          const { displayName, expiresIn, idToken, refreshToken, email } =
            response._tokenResponse;
          dispatch({
            type: ActionTypes.LOGIN,
            payload: {
              email: email,
              userName: displayName,
              token: response.user.accessToken,
              userId: idToken,
              error: null,
              isLoading: false,
            },
          });
          setCookie("token", response.user.accessToken, {
            maxAge: expiresIn,
          });
          setCookie("userId", data.userId, { maxAge: 60 * 60 * 24 });
          setCookie("refreshToken", refreshToken, {
            maxAge: 60 * 60 * 24 * 30,
          });
          await Navigate.push("/dashboard/exams");
        }
      }
    } catch (error: any) {
      toast.error(error.code.replace(/\W/g, " "));
      dispatch({
        type: ActionTypes.SIGN_UP,
        payload: {
          email: null,
          token: null,
          userName: null,
          userId: null,
          error: "",
          isLoading: false,
        },
      });
    }
  };
};
/*  ###########################################  find all exams action Creator  #####################################################  */
export const getAllExams = (token: string) => {
  return async (dispatch: Dispatch<Actions>) => {
    await refreshToken();
    try {
      const response = await axios.get(`${api}test/all`, {
        headers: {
          Authorization: token,
        },
      });
      const { data } = await response;
      dispatch({
        type: ActionTypes.GET_ALL_TESTS,
        payload: data,
      });
    } catch (err) {
      toast.error("we cannot get exams try agin");
    }
  };
};
/*  ###########################################  create exam Creator  #####################################################  */
export const createTest = (
  testId: string,
  _testType: string,
  _testDate: Date
) => {
  const raw = JSON.stringify({
    testDate: 0,
    testId: testId,
    testType: 0,
    userId: getCookie("userId"),
  });
  return async (dispatch: Dispatch<Actions>) => {
    try {
      const response = await axios.post(`${api}user-test`, raw, {
        headers: {
          Authorization: getCookie("token") as string,
          "content-type": "application/json",
        },
        method: "POST",
      });
      const { data } = await response;
      setCookie("testId", testId, { maxAge: 60 * 60 * 24 });
      toast.success("Your Exam successfully created");
      setTimeout(() => {
        Navigate.push({
          pathname: "/dashboard/exams/exam",
          query: { section: "reading", level: "intro" },
        });
      }, 3000);
      dispatch({
        type: ActionTypes.CREATE_TEST,
        payload: data,
      });
    } catch (err: any) {
      toast.error(err.response.data);
    }
  };
};
/*  ###########################################  get test Creator  #####################################################  */
export const getTest = (testId: string) => {
  return async (dispatch: Dispatch<Actions>) => {
    try {
      const response = await axios.get(`${api}test/full/${testId}`, {
        headers: {
          Authorization: getCookie("token") as string,
          "content-type": "application/json",
        },
        method: "POST",
      });
      const { data } = await response;
      setCookie("testId", testId, { maxAge: 60 * 60 * 24 });
      toast.success("Your Exam successfully created");
      setTimeout(() => {
        Navigate.push({
          pathname: "/dashboard/exams/exam",
          query: { section: "reading", level: "intro" },
        });
      }, 3000);
      dispatch({
        type: ActionTypes.GET_TEST,
        payload: data,
      });
    } catch (err: any) {
      toast.error(err.response.data);
    }
  };
};
/*  ########################################### answer 1 test Creator #####################################################  */

export const AnswerTest = (
  answer: string,
  questionId: string,
  userTestId: string
) => {
  const raw = JSON.stringify({
    answer,
    questionId,
    userTestId,
  });
  return async (dispatch: Dispatch<Actions>) => {
    try {
      const response = await axios.post(`${api}user-answer`, raw, {
        headers: {
          "content-type": "application/json",
          Authorization: getCookie("token") as string,
        },
        method: "POST",
      });
      console.log(response);
    } catch (err: any) {
      toast.error(err.response.data);
      console.log(err);
    }
  };
};
