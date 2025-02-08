
import girlWithToys from '../../../images/girl_with_toys.webp';
import './Home.css';

function Home() {
    return (
        <div className='home'>
            <div className='home images'>
                <div className='image1'>
                    {/* Text on top of the background image */}
                    <div className='text-overlay'>
                       Shop Your Favorite Toys
                       <button className='seeall'>See All</button>
                    </div>
                </div>
                <img className='image2' src={girlWithToys} alt="kids" />
            </div>
        </div>
    );
}

export default Home;
