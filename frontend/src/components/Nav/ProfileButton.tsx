import { useState, useEffect, useRef, MouseEvent } from 'react';
// import { useDispatch } from 'react-redux';

import OpenModalButton from "../OpenModalButton/OpenModalButton";
import LoginFormPage from "../LoginFormPage/LoginFormPage";
// import SignupFormModal from "../SignupFormModal/SignupFormModal"; // Ensure this is imported correctly
// import { useNavigate } from "react-router-dom";
import SignupFormModal from '../SignupFormModal/SignupFormModal';

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
//   const navigate = useNavigate();
//   const dispatch = useDispatch();
  const [showMenu, setShowMenu] = useState(false);
  const ulRef = useRef<HTMLDivElement | null>(null);

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

  const logout = (e: MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    // dispatch(sessionActions.logout());
    setShowMenu(false);
    // navigate('/');
  };

  const ulClassName = "profile-dropdown" + (showMenu ? "" : " hidden");

  return (
    <>
      <div className="profile" onClick={toggleMenu}>
        <button>Profile</button>
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
            </div>
          )}
        </div>
      )}
    </>
  );
}

export default ProfileButton;
