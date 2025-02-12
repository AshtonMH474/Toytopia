import { useDispatch, useSelector } from "react-redux"
import { AppDispatch, RootState } from "../../store/store"
import { useEffect } from "react"
import { getToys } from "../../store/toys"
import { Link } from "react-router-dom"
import './HighlyRated.css'

function HighlyRated() {
    const dispatch = useDispatch<AppDispatch>();
    const toys = useSelector((state: RootState) => state.toys.toys);
    // const [currIndex,setIndex] = useState<number>(0)

    useEffect(() => {
        async function filterToys() {
            await dispatch(getToys()); // Ensure we have toys from the state
        }

        filterToys();
    }, [dispatch]);

    // Filter the toys after the toys are loaded into the state
    const ratedToys = toys.filter((toy) => toy.rating >= 4.8);
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
                    {ratedToys && ratedToys.length ? (
                        ratedToys.slice(0, 0 + 4).map((toy) => (
                            <div key={toy.id} className="currToy">
                                <img className="img" src={toy.images[0].img_url}/>
                                <div className="card">

                                </div>
                                {/* <h2>{toy.company}</h2> */}

                            </div>
                        ))
                    ) : (
                        <p>No highly rated toys found.</p>
                    )}
                </div>
            </div>
        </>
    );
}

export default HighlyRated;
