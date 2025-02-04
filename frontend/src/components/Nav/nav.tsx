
// import { useSelector } from 'react-redux';
// import { IoMenu } from "react-icons/io5";
// import { IoIosExit } from "react-icons/io";
import dino from '../../../images/dino.png'
import { CgProfile } from "react-icons/cg";
import { MdToys } from "react-icons/md";
import { MdDashboardCustomize } from "react-icons/md";
import { MdReviews } from "react-icons/md";
import { FaHome } from "react-icons/fa";
// import ProfileButton from './ProfileButton';

// import { RootState } from '../../store/store';
// import { User } from '../../store/session';
import { Link } from 'react-router-dom';
import './nav.css';


function Navigation({isLoaded}) {
    // const sessionUser = useSelector<RootState, User | null>((state) => state.user.user);


    return (
        <>
        {isLoaded && (
        <div className='sidebar'>

            <div className='logoContainer'>
                <img src={dino} alt="dino" className='dinoLogo'/>
                <h2 className='title'>Toytopia</h2>
            </div>
            <div className='burgerContainer'>
                <div className='burgerTrigger'></div>
                <div className='burgerMenu'></div>
            </div>
            <div className='profileContainer'>
                <CgProfile className='imgProfile'/>
                <div className='profileContents'>
                    <p className='name'>Hello, John</p>
                    <p className='email'>johnsmith@gmail.com</p>

                </div>
            </div>
            <div className='contentsContainer'>
                <ul>
                    <li>
                        <Link to='/'>
                            <FaHome/>
                            Home
                        </Link>
                    </li>
                    <li>
                        <Link to='/toys'>
                            <MdToys/>
                            Toys
                        </Link>
                    </li>
                    <li>
                        <Link to='/wishlists'>
                            <MdDashboardCustomize/>
                            Wishlist
                        </Link>
                    </li>
                    <li>
                        <Link to='/reviews'>
                            <MdReviews/>
                            Reviews
                        </Link>
                    </li>
                </ul>
            </div>

        </div>
        )}
        </>
    )
//   return (
//     <div className='containerNav'>



//         {isLoaded && (
//              <div className='allProfile'>
//              <div>
//              {/* <ProfileButton user={sessionUser} /> */}

//            </div>
//          </div>
//         )}



//         </div>


//   );
}

export default Navigation;
