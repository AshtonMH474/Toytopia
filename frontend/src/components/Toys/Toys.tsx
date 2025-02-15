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

    const [visableThemes, setThemes] = useState<boolean>(false)
    const [visablePrices, setPrices] = useState<boolean>(false)
    const [visableBrands, setBrands] = useState<boolean>(false)

    const [theme,setTheme] = useState<string>("")
    const [minPrice,setMinPrice] = useState<number>(0)
    const [maxPrice, setMaxPrice] = useState<number>(Infinity)

    useEffect(() => {
        async function grabToys(){
            await dispatch(getToys())

        }
        grabToys()
    },[dispatch])


    const handleMax = (e: React.ChangeEvent<HTMLInputElement>) => {
        const value = e.target.value;
        // Parse the string input into a number, if possible
        setMaxPrice(value ? parseFloat(value) : 0);
      };
      const handleMin = (e: React.ChangeEvent<HTMLInputElement>) => {
        const value = e.target.value;
        // Parse the string input into a number, if possible
        setMinPrice(value ? parseFloat(value) : 0);
      };
    return (
        <>
        <div className="filterToys">
            <div className="companys">
                <h2 onClick={() => setBrands(!visableBrands)}>Brands <FaArrowDown/></h2>
                <Brands visable={visableBrands}/>
            </div>
            <div className="themes">
                <h2 onClick={() => setThemes(!visableThemes)}>Theme <FaArrowDown/></h2>
                <div className={`optionsThemes ${visableThemes? "" : "hideOptions"}`}>
                    <label>
                    Theme
                    <input type="search" onChange={(e) => setTheme(e.target.value)} value={theme}/>
                    </label>
                </div>
            </div>
            <div className="prices">
                <h2 onClick={() => setPrices(!visablePrices)}>Prices <FaArrowDown/></h2>
                <div className={`optionsPrice ${visablePrices? "" : "hideOptions"}`}>
                    <label>
                    Minumum Price
                    <input type="number" onChange={handleMin} min="0" value={minPrice}/>
                    </label>
                    <label>
                    Max Price
                    <input type="number" onChange={handleMax} min="0" value={maxPrice}/>
                    </label>
                </div>
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
