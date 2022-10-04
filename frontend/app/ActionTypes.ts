export enum ActionTypes {
  LOGIN = "LOGIN",
  SIGN_UP = "SIGN_UP",
  GET_ALL_TESTS = "GET_ALL_TESTS",
  CREATE_TEST = "CREATE_TEST",
  GET_TEST = "GET_TEST",
}

export type authState = {
  email: string | null;
  token: string | null;
  userName: string | null;
  error: string | null;
  userId: string | null;
  isLoading: boolean;
};
export type testsState = {
  tests:
    | {
        id: string;
        name: string;
        description: string;
        testDate: string;
        section: string;
      }[]
    | null;
  createdTest: {
    id: string;
    name: string;
    description: string;
    module: string;
    testDate: string;
    createdAt: string;
  };
  test: any;
};

export type ArticleAction = {
  type: string;
  payload: string[] | [] | {} | string | number;
};

export type StateType = authState | testsState;

export type DispatchType = (args: ArticleAction) => ArticleAction;
