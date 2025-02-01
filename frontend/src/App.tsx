
import './App.css'
import Navigation from './components/Nav/nav';
import { createBrowserRouter, Outlet, RouterProvider } from 'react-router-dom';
import Home from './components/Home/Home';
function Layout() {




  return (
    <>
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
  return <RouterProvider router={router} />;
}

export default App;
