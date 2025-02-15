
import { authFetch } from "./auth";
import { User } from "./session";
const SET_TOYS = 'toys/SET_TOYS'


export interface ToyImage{
    id:number;
    img_url:string
    primary_img:boolean
}
export interface Toy {
    id: number;
    release_date: Date
    price: number
    company: string
    product_type: string
    theme:string
    count: number
    available: boolean
    rating: number
    user_id: number
    user: User
    images:ToyImage[]
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

export const filterToys = (theme) => async (dispatch:React.Dispatch<ToyActions>) => {
    let url = `/api/toys?&`;
    if(theme && theme.length) url = url + `theme=${theme}&`


    const res = await authFetch(url)
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
