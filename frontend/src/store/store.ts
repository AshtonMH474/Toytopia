import { createStore, combineReducers, applyMiddleware, compose, Middleware,Store,StoreEnhancer } from 'redux';
import thunk from 'redux-thunk';

// Define your reducers
const rootReducer = combineReducers({});

let enhancer: StoreEnhancer;

if (import.meta.env.MODE === 'production') {
  enhancer = applyMiddleware(thunk as unknown as Middleware);
} else {
  // Dynamically import redux-logger for development
  const logger = (await import('redux-logger')).default;
  const composeEnhancers =
    (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

  enhancer = composeEnhancers(applyMiddleware(thunk as unknown as Middleware, logger));
}

interface RootState {
    // Your state properties here, e.g.:
    // user: UserState;
  }

  const configureStore = (preloadedState?: RootState): Store<RootState> => {
    return createStore(rootReducer, preloadedState, enhancer);
  };

  export default configureStore;
