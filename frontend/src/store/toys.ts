
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



export const filterToys = (filters?:any) => async (dispatch:React.Dispatch<ToyActions>) => {
    let url = `/api/toys?&`;
    if(filters){
        if(filters.theme && filters.theme.length) url = url + `theme=${filters.theme}&`
        if(filters.product && filters.product.length) url = url + `product_type=${filters.product}&`
        if(filters.minPrice)url = url + `min_price=${filters.minPrice}&`
        url = url + `max_price=${filters.maxPrice}&company=`

        for(let key in filters.brands){
            if(filters.brands[key] == true){
                url = url + `${key},`
            }
        }
}
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
