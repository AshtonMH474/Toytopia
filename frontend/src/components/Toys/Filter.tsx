import { useEffect, useState } from "react";
import { AppDispatch } from "../../store/store";
import { useDispatch } from "react-redux";
import Brands from "./Brands";
import { filterToys } from "../../store/toys";
import { FaArrowDown } from "react-icons/fa";
import { FaArrowRight } from "react-icons/fa";
import './Filters.css'


function FilteredToys(){
        const dispatch = useDispatch<AppDispatch>();
        const [visableThemes, setThemes] = useState<boolean>(false)
        const [visablePrices, setPrices] = useState<boolean>(false)
        const [visableBrands, setBrands] = useState<boolean>(false)
        const [visableProducts, setProducts] = useState<boolean>(false)


        const [theme,setTheme] = useState<string>("")
        const [product,setProduct] = useState<string>("")
        const [minPrice,setMinPrice] = useState<number>(0)
        const [maxPrice, setMaxPrice] = useState<number>(Infinity)
        const [brands, setObjBrands] = useState({
            Disney:false,
            Hasbro:false,
            PlaymatesToys:false,
            LEGO:false,
            Mattel:false,
            hotwheels:false
        })


        useEffect(() => {
            async function filterStuff() {
            // filtering themes
            let filters = {
                theme: theme.length? theme : null,
                product: product.length? product : null,
                minPrice:minPrice,
                maxPrice:maxPrice,
                brands:brands
            }

            await dispatch(filterToys(filters))

            }

        filterStuff()
    },[dispatch,theme,minPrice,maxPrice,brands,product])

    const handleMax = (e: React.ChangeEvent<HTMLInputElement>) => {
        const value = e.target.value;
        // Parse the string input into a number, if possible

        setMaxPrice(parseFloat(value));
      };
      const handleMin = (e: React.ChangeEvent<HTMLInputElement>) => {
        const value = e.target.value;
        // Parse the string input into a number, if possible
        if(!value)setMinPrice(0)
        setMinPrice(parseFloat(value));
      };
    return (
        <div className="filters-container">
                    <div className="products">

                        <div className="title">{!visableProducts && (<h2 onClick={() => setProducts(!visableProducts)}>Toy Name<FaArrowRight className="arrow"/></h2>)}</div>
                        <div className="title">{visableProducts && (<h2 onClick={() => setProducts(!visableProducts)}>Toy Name<FaArrowDown className="arrow"/></h2>)}</div>
                        <div className={`optionsThemes ${visableProducts? "" : "hideOptions"}`}>
                            <label>
                            Name
                            </label>
                            <input type="search" onChange={(e) => setProduct(e.target.value)} value={product}/>

                        </div>
                    </div>
                    <div className="companys">
                    {!visableBrands && (<h2 onClick={() => setBrands(!visableBrands)}>Brands<FaArrowRight className="arrow"/></h2>)}
                    {visableBrands && (<h2 onClick={() => setBrands(!visableBrands)}>Brands <FaArrowDown className="arrow"/></h2>)}
                        <Brands setBrands={setObjBrands} visable={visableBrands} />
                    </div>
                    <div className="themes">
                    {!visableThemes && (<h2 onClick={() => setThemes(!visableThemes)}>Themes<FaArrowRight className="arrow"/></h2>)}
                    {visableThemes && (<h2 onClick={() => setThemes(!visableThemes)}>Themes<FaArrowDown className="arrow"/></h2>)}
                        <div className={`optionsThemes ${visableThemes? "" : "hideOptions"}`}>
                            <label>
                            Theme
                            </label>
                            <input type="search" onChange={(e) => setTheme(e.target.value)} value={theme}/>

                        </div>
                    </div>
                    <div className="prices">
                        {!visablePrices && (<h2 onClick={() => setPrices(!visablePrices)}>Prices<FaArrowRight className="arrow"/></h2>)}
                        {visablePrices && (<h2 onClick={() => setPrices(!visablePrices)}>Prices<FaArrowDown className="arrow"/></h2>)}

                        <div className={`optionsPrice ${visablePrices? "" : "hideOptions"}`}>
                            <label>
                            Minumum Price
                            </label>
                            <input type="number" onChange={handleMin} min="0" value={minPrice}/>

                            <label>
                            Max Price
                            </label>
                            <input type="number" onChange={handleMax} min="0" value={maxPrice}/>

                        </div>
                    </div>
                    <button>Close Menu</button>
        </div>
    )
}

export default FilteredToys
