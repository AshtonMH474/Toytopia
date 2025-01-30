import { authFetch } from './auth';

// Action types
const SET_USER = 'session/setUser';
const REMOVE_USER = 'session/removeUser';

// Action interfaces
interface User {
  id: number;
  username: string;
  email: string;
}

interface SetUserAction {
  type: typeof SET_USER;
  payload: User;
}

interface RemoveUserAction {
  type: typeof REMOVE_USER;
}

export type SessionActions = SetUserAction | RemoveUserAction;

// Action creators
export const setUser = (user: User): SetUserAction => ({
  type: SET_USER,
  payload: user,
});

export const removeUser = (): RemoveUserAction => ({
  type: REMOVE_USER,
});

// Thunk action for login
interface LoginPayload {
  credential: string;
  password: string;
}

export const login = (user: LoginPayload) => async (dispatch: React.Dispatch<SessionActions>) => {
  const { credential, password } = user;
  const response = await authFetch('/api/login', {
    method: 'POST',
    body: JSON.stringify({ credential, password }),
  });
  const data = await response.json();
  dispatch(setUser(data.user));
  return response;
};

// Session state interface
export interface SessionState {
  user: User | null;
}

// Initial state
const initialState: SessionState = { user: null };

// Reducer
const sessionReducer = (state = initialState, action: SessionActions): SessionState => {
  switch (action.type) {
    case SET_USER:
      return { ...state, user: action.payload };
    case REMOVE_USER:
      return { ...state, user: null };
    default:
      return state;
  }
};

export default sessionReducer;
