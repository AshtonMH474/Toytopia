
import { useSelector } from 'react-redux';
import { useModal } from '../../Context/Modal';
import SignupFormModal from '../SignupFormModal/SignupFormModal';
import './Home.css';
import { RootState } from '../../store/store';
import { User } from '../../store/session';
import { useNavigate } from 'react-router-dom';
import Featured from './Featured';
import HighlyRated from './HighlyRated';

function Home() {
    const sessionUser = useSelector<RootState, User | null>((state) => state.user.user);
    const {setModalContent} = useModal()
    const nav = useNavigate()

    const signup = () => {
        if (!sessionUser) setModalContent(<SignupFormModal />);
        else nav('/wishlists')
    }

    return (
        <>
        <div className='homeStart'>
            <div className='home images'>
                <div onClick={() => nav('/toys')} className='image1'>
                    <div className='text-overlay'>
                       Shop Your Favorite Toys
                       <button className='seeall'>See All</button>
                    </div>
                </div>
                <div onClick={signup} className='image2'>
                {!sessionUser && (
                    <div className='text-overlay overlay2'>
                    Join Us Now
                     <button  className='seeall'>Sign Up</button>
                </div>
                )}
                {sessionUser && (
                    <div className='text-overlay overlay2'>
                    View My Wishlist
                     <button  className='seeall'>View</button>
                </div>
                )}

                </div>

            </div>

        </div>

        <div className='highlyrated'>
            <HighlyRated />
        </div>
        <div className='featured'>
            <Featured/>
        </div>
        </>

    );
}

export default Home;
