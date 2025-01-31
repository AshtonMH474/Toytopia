import { USER_FAILED, USER_LOGIN, USER_LOGOUT, USER_SUCCESS } from "../Constants/userConstants";

export interface UserState {
    loading?: boolean;
    error?: string | null | undefined;
    userInfo: { firstName?: string, lastName?: string, email?: string, username?: string, id?: number } | null;
}

interface Action {
    type: string;
    payload?: { firstName?: string, lastName?: string, email?: string, username?: string, id?: number } | null | string;
}

export const userLoginReducer = (state: UserState = { userInfo: null }, action: Action): UserState => {
    switch (action.type) {
        case USER_LOGOUT:
            return { userInfo: null };
        case USER_FAILED:
            // Ensure error is only set to a string or null
            return { loading: false, error: typeof action.payload === 'string' ? action.payload : null, userInfo: null };
        case USER_SUCCESS:
            // Ensure the payload is valid for userInfo
            const userInfo = action.payload && typeof action.payload !== 'string' ? action.payload : null;
            return { loading: false, userInfo };
        case USER_LOGIN:
            return { loading: true, userInfo: null };
        default:
            return state;
    }
};
