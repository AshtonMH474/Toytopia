import { AnyAction, combineReducers, createStore, StoreEnhancer } from "redux";
import { applyMiddleware, compose, } from "redux";
import {thunk, ThunkDispatch} from "redux-thunk";

import { Dispatch } from "react";
import sessionReducer from "./session";


// Define reducers
const reducers = combineReducers({
  // your reducers here
  user:sessionReducer
});

// Initial state
const initialState = {};

// Define the enhancer type as StoreEnhancer
let enhancer: StoreEnhancer<any, any> | undefined;

// if (import.meta.env.MODE === "production") {
if (import.meta.env.PROD){
  // In production, apply only the thunk middleware
  enhancer = applyMiddleware(thunk as any);
} else {
  // In development, dynamically import redux-logger
  const logger = (await import("redux-logger")).default;

  // Type assertion for window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__
  const composeEnhancers =
    (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

  // Apply both thunk and logger middleware
  enhancer = composeEnhancers(applyMiddleware(thunk as any, logger));
}

// Create the Redux store
const store = createStore(
  reducers,
  initialState,
  enhancer
);

export default store;
export type RootState = ReturnType<typeof reducers>
export type AppDispatch = Dispatch<AnyAction> & ThunkDispatch<RootState, unknown, AnyAction>;
