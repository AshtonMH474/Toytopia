import { useRef, useState, useContext, createContext, ReactNode } from 'react';
import ReactDOM from 'react-dom';
import './Modal.css';

// Define the modal context type
interface ModalContextType {
  modalRef: React.RefObject<HTMLDivElement>;
  modalContent: ReactNode | null;
  setModalContent: (content: ReactNode | null) => void;
  setOnModalClose: (callback: (() => void) | null) => void;
  closeModal: () => void;
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

  const closeModal = () => {
    setModalContent(null);
    if (onModalClose) {
      setOnModalClose(null);
      onModalClose();
    }
  };

  const contextValue: ModalContextType = {
    modalRef,
    modalContent,
    setModalContent,
    setOnModalClose,
    closeModal,
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
