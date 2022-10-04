import { ActionTypes } from "app/ActionTypes";

interface Login {
  type: ActionTypes.LOGIN;
  payload: {
    email: string | null;
    token: string | null;
    userId: string | null;
    userName: string | null;
    isLoading: boolean;
    error: string | null;
  };
}

interface SignUp {
  type: ActionTypes.SIGN_UP;
  payload: {
    email: string | null;
    userName: string | null;
    token: string | null;
    userId: string | null;
    isLoading: boolean;
    error: string | null;
  };
}

interface getAllTest {
  type: ActionTypes.GET_ALL_TESTS;
  payload: {
    id: string;
    name: string;
    description: string;
    testDate: string;
    section: string;
  }[];
}

interface createTest {
  type: ActionTypes.CREATE_TEST;
  payload: {
    id: string;
    name: string;
    description: string;
    module: string;
    testDate: string;
    createdAt: string;
  };
}

interface getTest {
  type: ActionTypes.GET_TEST;
  payload: any;
}

export type Actions = Login | SignUp | getAllTest | createTest | getTest;
