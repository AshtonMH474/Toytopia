// import './App.css'
import Navigation from './components/Nav/nav';
import { createBrowserRouter, Outlet, RouterProvider } from 'react-router-dom';
import Home from './components/Home/Home';
import { useDispatch } from 'react-redux';
import { useEffect, useState } from 'react';
import { restoreUser } from './store/session';
import { AppDispatch } from './store/store';
function Layout() {
  const dispatch = useDispatch<AppDispatch>();
  const [isLoaded, setIsLoaded] = useState<boolean>(false);

  useEffect(() => {
    dispatch(restoreUser()).then(() => {
    console.log("User restored");
    setIsLoaded(true);
  }).catch(error => {
    console.error("Error restoring user:", error);
  });
  }, [dispatch]);

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
      }
    ]
  }
]);

function App() {
  return <RouterProvider router={router} />;
}

export default App;
