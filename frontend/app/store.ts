import {applyMiddleware, combineReducers, createStore, Store} from "redux";
import {composeWithDevTools} from "redux-devtools-extension";
import thunk from "redux-thunk";
import {DispatchType} from "app/ActionTypes";
import authReducer from "app/authReducer";
import TestReducer from "./testsReducer";

export const store: Store & {
    dispatch: DispatchType;
} = createStore(
    combineReducers({
        auth: authReducer,
        tests: TestReducer
    }),
    composeWithDevTools(applyMiddleware(thunk))
);

export type RootState = ReturnType<typeof store.getState>;
