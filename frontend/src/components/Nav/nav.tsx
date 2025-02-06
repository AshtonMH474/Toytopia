
import { useDispatch, useSelector } from 'react-redux';
import dino from '../../../images/dino.png'
import { CgProfile } from "react-icons/cg";
import { MdToys } from "react-icons/md";
import { MdDashboardCustomize } from "react-icons/md";
import { MdReviews } from "react-icons/md";
import { FaHome } from "react-icons/fa";
import { IoIosExit } from "react-icons/io";
import { AppDispatch, RootState } from '../../store/store';
import { logoutUser, User } from '../../store/session';
import { Link } from 'react-router-dom';
import './nav.css';
import { useState } from 'react';
import LoginFormPage from '../LoginFormPage/LoginFormPage';
import SignupFormModal from '../SignupFormModal/SignupFormModal';
import { useModal } from '../../Context/Modal';
import { MouseEvent } from 'react';


function Navigation({isLoaded}) {
    const sessionUser = useSelector<RootState, User | null>((state) => state.user.user);
    const [showMenu,setMenu] = useState<boolean>(false)
    const {setModalContent,closeModal} = useModal()
    const dispatch = useDispatch<AppDispatch>()
    console.log(sessionUser)

    const handleCloseMenu = () => {
        setMenu(!showMenu)
    }

    const signup = () => {
        setModalContent(<SignupFormModal />);
    }
    const login = () => {
        setModalContent(<LoginFormPage />);
    }
   const logout = async (e: MouseEvent<HTMLElement>) => {
       e.preventDefault();
       await dispatch(logoutUser());
       await closeModal()
     };

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
                {sessionUser ? (<div className='profileContents'>
                    <p className='name'>Hello, {sessionUser.first_name}</p>
                    <p className='email'>{sessionUser.email}</p>

                </div>):(
                    <div className='profileContents'>
                        <p onClick={login} className='login options'>Login</p>

                        <p onClick={signup} className='signup options'>Sign Up</p>
                    </div>
                )}

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
                    {sessionUser && (<li>
                        <Link onClick={logout} to = '/' className='link'>
                            <IoIosExit className='logo'/>
                            <div>Logout</div>
                        </Link>
                    </li>)}
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
