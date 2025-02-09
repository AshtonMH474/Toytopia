
import { useModal } from '../../Context/Modal';
import SignupFormModal from '../SignupFormModal/SignupFormModal';
import './Home.css';

function Home() {
    const {setModalContent} = useModal()

    const signup = () => {
        setModalContent(<SignupFormModal />);
    }

    return (
        <div className='home'>
            <div className='home images'>
                <div className='image1'>
                    <div className='text-overlay'>
                       Shop Your Favorite Toys
                       <button className='seeall'>See All</button>
                    </div>
                </div>
                <div className='image2'>
                <div className='text-overlay overlay2'>
                        Join Us Now
                       <button onClick={signup} className='seeall'>Sign Up</button>
                    </div>
                </div>

            </div>
        </div>
    );
}

export default Home;
