import { useDispatch, useSelector } from "react-redux"
import { AppDispatch, RootState } from "../../store/store"
import { useEffect , useState} from "react"
import { getToys } from "../../store/toys"
import { HiArrowSmallRight } from "react-icons/hi2";
import { HiArrowSmallLeft } from "react-icons/hi2";
import { Link } from "react-router-dom"
import { IoStarSharp } from "react-icons/io5";
import './HighlyRated.css'

function HighlyRated() {
    const dispatch = useDispatch<AppDispatch>();
    const toys = useSelector((state: RootState) => state.toys.toys);
    const [currIndex,setIndex] = useState<number>(0)

    useEffect(() => {
        async function filterToys() {
            await dispatch(getToys()); // Ensure we have toys from the state
        }

        filterToys();
    }, [dispatch]);

    // Filter the toys after the toys are loaded into the state
    const ratedToys = toys.filter((toy) => toy.rating >= 4.8);


     function nextToys(){
     setIndex(currIndex + 4)
    }

    function prevToys(){
        if(currIndex > 3) setIndex(currIndex - 4)
        else setIndex(0)
    }


    console.log(ratedToys)
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
                    <button onClick={prevToys}><HiArrowSmallLeft /></button>
                    {ratedToys && ratedToys.length ? (
                        ratedToys.slice(currIndex, currIndex + 4).map((toy) => (
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
                                </div>


                            </div>
                        ))
                    ) : (
                        <p>No highly rated toys found.</p>
                    )}
                    <button onClick={nextToys}><HiArrowSmallRight/></button>
                </div>
            </div>
        </>
    );
}

export default HighlyRated;
