import { ThunkAction, ThunkDispatch } from "redux-thunk";
import {USER_FAILED, USER_SUCCESS } from "../Constants/userConstants";
import { AnyAction } from "redux";
import { authFetch } from "../store/auth";

export const login = (email,password) : ThunkAction<Promise<void>,RootState,unknown,AnyAction> =>
    async (dispatch: ThunkDispatch<RootState,unknown,AnyAction>): Promise<void>=> {
    try{
        dispatch({
            type:USER_SUCCESS
        })

        const res =  await authFetch('/api/login', {
            method: 'POST',
            body: JSON.stringify({ email, password }),
        })
        const data = await res.json();
        const userData = {id:data.user.id,firstName:data.user.first_name,lastName:data.user.last_name,email:data.user.email,username:data.user.username}
        dispatch({
            type:USER_SUCCESS,
            payload:userData
        })
    }catch(e :unknown) {
        const errorMessage =
        e instanceof Error ? e.message : 'An unknown error occurred';

        dispatch({
            type:USER_FAILED,
            payload: errorMessage
        })
    }

};
