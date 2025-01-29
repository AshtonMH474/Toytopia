import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { Provider } from 'react-redux';
import './index.css'
import App from './App.tsx'
import configureStore from './store/store.js';
import { Modal, ModalProvider } from './Context/Modal.tsx';


const store = configureStore();

if (process.env.NODE_ENV !== 'production') {
  window.store = store;
}

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ModalProvider>
    <Provider store={store}>
      <App />
      <Modal/>
    </Provider>
    </ModalProvider>
  </StrictMode>,
)
