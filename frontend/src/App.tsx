// import './App.css'
import Navigation from './components/Nav/nav';
import { createBrowserRouter, Outlet, RouterProvider, useLocation } from 'react-router-dom';
import Home from './components/Home/Home';
import { useDispatch } from 'react-redux';
import { useEffect, useState } from 'react';
import { restoreUser } from './store/session';
import { AppDispatch } from './store/store';
import Toys from './components/Toys';
import Wishlists from './components/Wishlist';
import Reviews from './components/Reviews';
import { useModal } from './Context/Modal';
function Layout() {
  const dispatch = useDispatch<AppDispatch>();
  const [isLoaded, setIsLoaded] = useState<boolean>(false);
  const location = useLocation();
  const {setObjBrands} = useModal()

  useEffect(() => {
    if(location.pathname != '/toys'){
      setObjBrands({
        Disney: false,
        Hasbro: false,
        PlaymatesToys: false,
        LEGO: false,
        Mattel: false,
        Hotwheels: false
    })
    }
    window.scrollTo(0, 0);
    dispatch(restoreUser()).then(() => {
    console.log("User restored");
    setIsLoaded(true);
  }).catch(error => {
    console.error("Error restoring user:", error);
  });
  }, [dispatch,location]);

  return (
    <>
      <Navigation isLoaded={isLoaded}/>
      {isLoaded && <Outlet />}
    </>
  );
}

const router = createBrowserRouter([
  {
    element: <Layout />,
    children: [
      {
        path: '/',
        element: <Home/>
      },
      {
        path:'/toys',
        element:<Toys/>
      },
      {
        path:'/wishlists',
        element:<Wishlists/>
      },
      {
        path:'/reviews',
        element:<Reviews/>
      }
    ]
  }
]);

function App() {
  return <RouterProvider router={router} />;
}

export default App;
