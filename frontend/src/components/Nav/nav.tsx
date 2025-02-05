
// import { useSelector } from 'react-redux';
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
import { useState } from 'react';


function Navigation({isLoaded}) {
    // const sessionUser = useSelector<RootState, User | null>((state) => state.user.user);
    const [showMenu,setMenu] = useState<boolean>(false)

    const handleCloseMenu = () => {
        setMenu(!showMenu)
    }
    return (
        <>
        {isLoaded && (
        <div className={showMenu == false ? 'sidebar' : "sidebar active"}>

            <div className={showMenu == false ? 'logoContainer' : 'logoContainer active'} >
                <img src={dino} alt="dino" className='dinoLogo'/>
                <h2 className='title'>Toytopia</h2>
            </div>
            <div className={showMenu == false ? 'burgerContainer' : 'burgerContainer active'}>
                <div onClick={handleCloseMenu} className='burgerTrigger'></div>
                <div className='burgerMenu'></div>
            </div>
            <div className={showMenu == false ? 'profileContainer' : 'profileContainer active'}>
                <div><CgProfile className='imgProfile'/></div>
                <div className='profileContents'>
                    <p className='name'>Hello, John</p>
                    <p className='email'>johnsmith@gmail.com</p>

                </div>
            </div>
            <div className={showMenu == false ? 'contentsContainer' : 'contentsContainer active'}>
                <ul>
                    <li>
                        <Link to='/' className='link'>
                            <FaHome className='logo'/>
                            <div>Home</div>
                        </Link>
                    </li>
                    <li>
                        <Link to='/toys' className='link'>
                            <MdToys className='logo'/>
                            <div>Toys</div>
                        </Link>
                    </li>
                    <li>
                        <Link to='/wishlists' className='link'>
                            <MdDashboardCustomize className='logo'/>
                            <div>Wishlist</div>
                        </Link>
                    </li>
                    <li>
                        <Link to='/reviews' className='link'>
                            <MdReviews className='logo'/>
                            <div>Reviews</div>
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
