import { NavLink } from 'react-router-dom';
import { useSelector } from 'react-redux';
import ProfileButton from './ProfileButton';
import { UserState } from '../../Reducers/userReducers';
import { RootState } from '../../store/store';
// import './Navigation.css';


function Navigation( ) {

    const sessionUser = useSelector<RootState, UserState>((state) => state.user);
    const {userInfo} = sessionUser

  return (
    <div className='containerNav'>




        <div className='allProfile'>
            <div>
            <ProfileButton user={userInfo ?? { firstName: '', lastName: '', email: '', username: '', id: null }} />

          </div>
        </div>


        </div>


  );
}

export default Navigation;
