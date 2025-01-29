
import { useEffect } from 'react';
import './App.css'
import Navigation from './components/Nav/nav';
import { createBrowserRouter, Outlet, RouterProvider } from 'react-router-dom';
import Home from './components/Home/Home';
function Layout() {




  return (
    <>
      <Navigation />
      <Outlet />
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

    useEffect(() => {
      document.title = "ashton's airbnb"

    }, [])
  return <RouterProvider router={router} />;
}

export default App;
