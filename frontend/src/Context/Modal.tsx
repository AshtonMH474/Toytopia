import { useRef, useState, useContext, createContext, ReactNode } from 'react';
import ReactDOM from 'react-dom';
import './Modal.css';

interface Brands {
  Disney: boolean;
  Hasbro: boolean;
  PlaymatesToys: boolean;
  LEGO: boolean;
  Mattel: boolean;
  Hotwheels: boolean;
}

// Define the modal context type
interface ModalContextType {
  modalRef: React.RefObject<HTMLDivElement>;
  modalContent: ReactNode | null;
  setModalContent: (content: ReactNode | null) => void;
  setOnModalClose: (callback: (() => void) | null) => void;
  closeModal: () => void;
  brands: Brands; // Use the specific Brands type
  setObjBrands: React.Dispatch<React.SetStateAction<Brands>>; // Add setObjBrands to the context
}



// Create the context with an initial undefined state
const ModalContext = createContext<ModalContextType | undefined>(undefined);

interface ModalProviderProps {
  children: ReactNode;
}

export function ModalProvider({ children }: ModalProviderProps) {
  const modalRef = useRef<HTMLDivElement>(null);
  const [modalContent, setModalContent] = useState<ReactNode | null>(null);
  const [onModalClose, setOnModalClose] = useState<(() => void) | null>(null);

  // Define the brands state
  const [brands, setObjBrands] = useState({
    Disney: false,
    Hasbro: false,
    PlaymatesToys: false,
    LEGO: false,
    Mattel: false,
    Hotwheels: false,
  });

  const closeModal = () => {
    setModalContent(null);
    if (onModalClose) {
      setOnModalClose(null);
      onModalClose();
    }
  };

  // Context value including brands and setObjBrands
  const contextValue: ModalContextType = {
    modalRef,
    modalContent,
    setModalContent,
    setOnModalClose,
    closeModal,
    brands, // Include brands in context
    setObjBrands, // Include setObjBrands in context
  };

  return (
    <>
      <ModalContext.Provider value={contextValue}>{children}</ModalContext.Provider>
      <div ref={modalRef} />
    </>
  );
}

export function Modal() {
  const modalContext = useContext(ModalContext);

  if (!modalContext) {
    throw new Error('Modal must be used within a ModalProvider');
  }

  const { modalRef, modalContent, closeModal } = modalContext;

  if (!modalRef.current || !modalContent) return null;

  return ReactDOM.createPortal(
    <div id="modal">
      <div id="modal-background" onClick={closeModal} />
      <div id="modal-content">{modalContent}</div>
    </div>,
    modalRef.current
  );
}

// Custom hook to use modal context
export const useModal = (): ModalContextType => {
  const context = useContext(ModalContext);
  if (!context) {
    throw new Error('useModal must be used within a ModalProvider');
  }
  return context;
};
