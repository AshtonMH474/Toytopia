import { useDispatch, useSelector } from "react-redux"
import { AppDispatch, RootState } from "../../store/store"
import { useEffect, useState } from "react"
import { getToys } from "../../store/toys"
import { Link } from "react-router-dom"
import './HighlyRated.css'

function HighlyRated(){
    const dispatch = useDispatch<AppDispatch>()
    const toys = useSelector((state: RootState) => state.toys.toys);
    const [ratedToys, setToys] = useState(toys)
    console.log(ratedToys)
    useEffect(() => {
        async function filterToys(){
            await dispatch(getToys())
            const newArr = toys.filter((toy) => toy.rating >= 4.8);
            await setToys(newArr)
        }

        filterToys()

    },[dispatch])
    return (
        <>
        <div>
            <div className="info">
                <h1>Fan Favorites</h1>
                <Link to='/toys' className='link'>
                    See All
                </Link>
            </div>
            <div className="toys">
                {ratedToys.length && ratedToys.map((toy) => (
                    <div>
                        <h2>{toy.company}</h2>
                    </div>
                ))}
            </div>
        </div>
        </>
    )
}


export default HighlyRated
