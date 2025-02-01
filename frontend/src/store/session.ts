import { authFetch } from './auth';

// Action types
const SET_USER = 'session/setUser';
const REMOVE_USER = 'session/removeUser';

// Action interfaces
export interface User {
  id: number;
  username: string;
  email: string;
  firstName:string;
  lastName:string
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
  email: string;
  password: string;
}

export const restoreUser = () => async (dispatch: React.Dispatch<SessionActions>) =>{
    try {
        // Make the API request
        const response = await authFetch('/api/users/current');

        // Parse the response
        const data = await response.json();

        // Log the entire data object to the console
        console.log("User data from API:", data);

        // Dispatch the action to store the user data
        await dispatch(setUser(data.user));

        // Return the response (optional)
        return response;
    } catch (error) {
        console.error("Error fetching user data:", error);
    }
}

export const login = (user: LoginPayload) => async (dispatch: React.Dispatch<SessionActions>) => {
  const { email, password } = user;
  const response = await authFetch('/api/login', {
    method: 'POST',
    body: JSON.stringify({ email:email, password }),
  });
  const data = await response.json();
  await dispatch(setUser(data.user));
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
