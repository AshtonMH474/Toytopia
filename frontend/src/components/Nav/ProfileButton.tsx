import { useState, useEffect, useRef, MouseEvent } from 'react';
import { useDispatch } from 'react-redux';
import OpenModalButton from "../OpenModalButton/OpenModalButton";
import LoginFormPage from "../LoginFormPage/LoginFormPage";
import SignupFormModal from '../SignupFormModal/SignupFormModal';
import { logoutUser } from '../../store/session';
import { useModal } from '../../Context/Modal';
import { AppDispatch } from '../../store/store';
import { IoMenu } from "react-icons/io5";
import { IoIosExit } from "react-icons/io";
import { useNavigate } from 'react-router-dom';


// Define user prop type
interface User {
  firstName: string;
  email: string;
}

// Define component props
interface ProfileButtonProps {
  user: User | null;
}

function ProfileButton({ user }: ProfileButtonProps) {
    const nav = useNavigate()
  const dispatch = useDispatch<AppDispatch>();
  const [showMenu, setShowMenu] = useState(false);
  const ulRef = useRef<HTMLDivElement | null>(null);
  const {closeModal} = useModal()

  const toggleMenu = (e: MouseEvent<HTMLDivElement>) => {
    e.stopPropagation();
    setShowMenu((prev) => !prev);
  };

  const closeMenu = (e: MouseEvent<Document>) => {
    if (ulRef.current && !ulRef.current.contains(e.target as Node)) {
      setShowMenu(false);
    }
  };

  useEffect(() => {
    if (!showMenu) return;

    document.addEventListener("click", closeMenu as any);

    return () => document.removeEventListener("click", closeMenu as any);
  }, [showMenu]);

  const logout = async (e: MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    await dispatch(logoutUser());
    setShowMenu(false);
    await closeModal()

  };
  const goHome = async() => {
    nav('/')
  }

  const ulClassName = "profile-dropdown" + (showMenu ? "" : " hidden");

  return (
    <>
      <div className="profile" onClick={toggleMenu}>
        <IoMenu/>
      </div>

      {showMenu && (
        <div className={ulClassName} ref={ulRef}>
          {user ? (
            <div className="boxProfile">
              <div>Hello, {user.firstName}</div>
              <div>{user.email}</div>
              <div>
                <button className="buttonProfile" onClick={logout}>Log Out</button>
              </div>
            </div>
          ) : (
            <div className="boxProfile">
                <div onClick={goHome}>
                    Home
                </div>
              <div>
                <OpenModalButton
                  buttonText="Log In"
                  modalComponent={<LoginFormPage />}
                />
              </div>

              <div>
                <OpenModalButton
                  buttonText="Sign Up"
                  modalComponent={<SignupFormModal />}
                />
              </div>

              <div onClick={toggleMenu}>
              <IoIosExit/>
              </div>

            </div>

          )}
        </div>
      )}
    </>
  );
}

export default ProfileButton;
