import { NavLink } from 'react-router-dom';
import { useSelector } from 'react-redux';
import ProfileButton from './ProfileButton';
// import './Navigation.css';


function Navigation( ) {

//   const sessionUser = useSelector((state:RootState)=> state.session.user);


  return (
    <div className='containerNav'>




        <div className='allProfile'>
            <div>
          {/* <ProfileButton user={sessionUser ?? null}  /> */}
          </div>
        </div>


        </div>


  );
}

export default Navigation;
