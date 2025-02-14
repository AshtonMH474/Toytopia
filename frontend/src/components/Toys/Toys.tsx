import { useDispatch, useSelector } from "react-redux";
import { AppDispatch, RootState } from "../../store/store"
import { useEffect, useState } from "react";
import { getToys } from "../../store/toys";
import { IoStarSharp } from "react-icons/io5";
import { FaArrowDown } from "react-icons/fa";
import './toys.css'
import Brands from "./Brands";

function Toys(){
    const dispatch = useDispatch<AppDispatch>();
    const toys = useSelector((state: RootState) => state.toys.toys);


    const [theme,setTheme] = useState<string>("")
    const [visableBrands, setBrands] = useState<boolean>(false)

    useEffect(() => {
        async function grabToys(){
            await dispatch(getToys())

        }
        grabToys()
    },[dispatch])
    return (
        <>
        <div className="filterToys">
            <div className="companys">
                <h2 onClick={() => setBrands(!visableBrands)}>Brands <FaArrowDown/></h2>
                <Brands visable={visableBrands}/>
            </div>
            <div>
            <h2 onClick={() => setBrands(!visableBrands)}>Theme <FaArrowDown/></h2>
            <label>
            Theme
            <input type="search" onChange={(e) => setTheme(e.target.value)} value={theme}/>
            </label>
            </div>

        </div>
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
