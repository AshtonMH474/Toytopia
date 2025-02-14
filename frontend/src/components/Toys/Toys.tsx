import { useDispatch, useSelector } from "react-redux";
import { AppDispatch, RootState } from "../../store/store"
import { useEffect } from "react";
import { getToys } from "../../store/toys";
import { IoStarSharp } from "react-icons/io5";
import './toys.css'

function Toys(){
    const dispatch = useDispatch<AppDispatch>();
    const toys = useSelector((state: RootState) => state.toys.toys);

    useEffect(() => {
        async function grabToys(){
            await dispatch(getToys())

        }
        grabToys()
    },[dispatch])
    return (
        <>
        <div className="toys">
            {toys && toys.length ? (
                toys.map((toy) => (
                    <div key={toy.id} className="currToy">
                        <div className="divImg"><img className="img" src={toy.images[0].img_url}/></div>
                        <div className="card">
                            <div className="Info">
                                <div className="rating">
                                    <IoStarSharp className={toy.rating >= 1 ? 'red' : 'gray'}/>
                                    <IoStarSharp className={toy.rating >= 2 ? 'red' : 'gray'}/>
                                    <IoStarSharp className={toy.rating >= 3 ? 'red' : 'gray'}/>
                                    <IoStarSharp className={toy.rating >= 4 ? 'red' : 'gray'}/>
                                    <IoStarSharp className={toy.rating >= 5 ? 'red' : 'gray'}/>
                                    {toy.rating.toFixed(1)}
                                    ({toy.count})
                                </div>
                                <div className="price">
                                    ${toy.price.toFixed(2)}
                                </div>
                            </div>
                            <div className="bottomInfo">
                                <div className="name">{toy.product_type}</div>
                                <button className="button">Add to Cart</button>
                            </div>
                        </div>
                    </div>
                ))
            ):(
               <div>Loading...</div>
            )}
        </div>
        </>
    )
}

export default Toys
