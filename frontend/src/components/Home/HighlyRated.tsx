import { useDispatch, useSelector } from "react-redux"
import { AppDispatch, RootState } from "../../store/store"
import { useEffect } from "react"
import { getToys } from "../../store/toys"

function HighlyRated(){
    const dispatch = useDispatch<AppDispatch>()
    const toys = useSelector((state: RootState) => state.toys.toys);
    console.log(toys)
    useEffect(() => {
        dispatch(getToys())
    },[dispatch])
    return (
        <>
        </>
    )
}


export default HighlyRated
