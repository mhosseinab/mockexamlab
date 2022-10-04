import { ActionTypes, StateType } from "app/ActionTypes";
import { Actions } from "app/Actions";

const initialState: StateType = {
  email: "",
  token: "",
  userId: "",
  userName: "sina soleymanzadeh",
  error: null,
  isLoading: false,
};

const AuthReducer = (
  state: StateType = initialState,
  action: Actions
): StateType => {
  switch (action.type) {
    case ActionTypes.LOGIN:
      return {
        ...state,
        email: action.payload.email,
        userId: action.payload.userId,
        userName: action.payload.userName,
        token: action.payload.token,
        error: action.payload.error,
        isLoading: action.payload.isLoading,
      };
    case ActionTypes.SIGN_UP:
      return {
        ...state,
        email: action.payload.email,
        userName: action.payload.userName,
        userId: action.payload.userId,
        token: action.payload.token,
        error: action.payload.error,
        isLoading: action.payload.isLoading,
      };
    default:
      return state;
  }
};

export default AuthReducer;
