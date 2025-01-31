import { USER_FAILED, USER_LOGIN, USER_LOGOUT, USER_SUCCESS } from "../Constants/userConstants";

export interface UserState {
    loading?:boolean,
    error?:string,
    userInfo:{firstName?:string, lastName?:string,email?:string,username?:string,id?:number}
}

interface Action {
    type:string,
    payload?:string
}

export const userLoginReducer = (state:UserState = {userInfo:{}},action:Action) => {
    switch(action.type){
        case USER_LOGOUT:
            return {userInfo:null}
        case USER_FAILED:
            return {loading:false, error:action.payload,userInfo:null}
        case USER_SUCCESS:
            return {loading:false,userInfo:action.payload}
        case USER_LOGIN:
            return {loading:true}
        default:
            return state

    }

}
