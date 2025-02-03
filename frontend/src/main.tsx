import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { Provider } from 'react-redux';
// import './index.css'
import App from './App.tsx'
import store from './store/store.tsx';
import { Modal, ModalProvider } from './Context/Modal.tsx';




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
