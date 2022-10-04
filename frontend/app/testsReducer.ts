import { ActionTypes, StateType } from "app/ActionTypes";

import { Actions } from "app/Actions";

const initialState: StateType = {
  tests: null,
  createdTest: {
    id: "",
    testDate: "",
    module: "",
    createdAt: "",
    description: "",
    name: "",
  },
  test: {},
};

const TestReducer = (
  state: StateType = initialState,
  action: Actions
): StateType => {
  switch (action.type) {
    case ActionTypes.GET_ALL_TESTS:
      return {
        ...state,
        tests: action.payload,
      };
    case ActionTypes.CREATE_TEST:
      return {
        ...state,
        createdTest: action.payload,
      };
    case ActionTypes.GET_TEST:
      return {
        ...state,
        test: action.payload,
      };
    default:
      return state;
  }
};

export default TestReducer;
