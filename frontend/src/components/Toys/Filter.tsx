import { useEffect, useState } from "react";
import { AppDispatch } from "../../store/store";
import { useDispatch } from "react-redux";
import Brands from "./Brands";
import { filterToys } from "../../store/toys";
import { FaArrowDown } from "react-icons/fa";


function FilteredToys(){
        const dispatch = useDispatch<AppDispatch>();
        const [visableThemes, setThemes] = useState<boolean>(false)
        const [visablePrices, setPrices] = useState<boolean>(false)
        const [visableBrands, setBrands] = useState<boolean>(false)

        const [theme,setTheme] = useState<string>("")
        const [minPrice,setMinPrice] = useState<number>(0)
        const [maxPrice, setMaxPrice] = useState<number>(Infinity)
        const [brands, setObjBrands] = useState({
            disney:false,
            hasbro:false,
            playmatesToys:false,
            lego:false,
            mattel:false,
            hotwheels:false
        })


        useEffect(() => {
            async function filterStuff() {
            // filtering themes
            let filters = {
                theme: theme.length? theme : null,
                minPrice:minPrice,
                maxPrice:maxPrice,
                brands:brands
            }


            await dispatch(filterToys(filters))

            }

        filterStuff()
    },[dispatch,theme,minPrice,maxPrice,brands])

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
        <>
        <div className="companys">
                        <h2 onClick={() => setBrands(!visableBrands)}>Brands <FaArrowDown/></h2>
                        <Brands setBrands={setObjBrands} visable={visableBrands} />
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
        </>
    )
}

export default FilteredToys
