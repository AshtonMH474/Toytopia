
import { authFetch } from "./auth";
import { User } from "./session";
const SET_TOYS = 'toys/SET_TOYS'


export interface ToyImage{
    id:number;
    ImgUrl:string
    PrimaryImage:boolean
}
export interface Toy {
    id: number;
    ReleaseDate: Date
    Price: number
    Company: string
    ProductType: string
    Theme:string
    Count: number
    Available: boolean
    Rating: number
    UserId: number
    User: User
    Images:ToyImage[]
}


interface SetToyAction{
    type: typeof SET_TOYS
    payload: Toy[]
}

export type ToyActions = SetToyAction

export const setToys = (toys:Toy[]): SetToyAction => ({
    type:SET_TOYS,
    payload:toys
})


export const getToys = () => async (dispatch: React.Dispatch<ToyActions>) => {
    const res = await authFetch(`/api/toys`)
    const data = await res.json()
    await dispatch(setToys(data))
    return res
}

export interface ToyState {
    toys:Toy[];
}
const initialState: ToyState = {toys:[]}

const toysReducer = (state = initialState, action:ToyActions ) : ToyState => {
    switch(action.type){
        case SET_TOYS:
            return {...state, toys:action.payload}
        default:
            return state;
    }
}

export default toysReducer
