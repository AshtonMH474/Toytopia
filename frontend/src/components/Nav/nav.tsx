
import { useSelector } from 'react-redux';
import ProfileButton from './ProfileButton';

import { RootState } from '../../store/store';
import { User } from '../../store/session';
// import './Navigation.css';


function Navigation( ) {

    const sessionUser = useSelector<RootState, User | null>((state) => state.user.user);


  return (
    <div className='containerNav'>




        <div className='allProfile'>
            <div>
            <ProfileButton user={sessionUser} />

          </div>
        </div>


        </div>


  );
}

export default Navigation;
