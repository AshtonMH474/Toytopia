import { useModal } from "../../Context/Modal";

interface OpenModalButtonProps {
  modalComponent: JSX.Element;
  buttonText: string;
  onButtonClick?: () => void; // Make optional
  onModalClose?: () => void; // Make optional
}

function OpenModalButton({
  modalComponent,
  buttonText,
  onButtonClick,
  onModalClose,
}: OpenModalButtonProps) {
  const { setModalContent, setOnModalClose } = useModal();

  const onClick = () => {
    if (onModalClose) setOnModalClose(onModalClose);
    setModalContent(modalComponent);
    if (onButtonClick) onButtonClick(); // Only call if it exists
  };

  if (buttonText === "Update")
    return <button className="deleteModal" onClick={onClick}>{buttonText}</button>;
  if (buttonText === "Delete")
    return <button className="deleteModal" onClick={onClick}>{buttonText}</button>;
  if (buttonText === "Create your Review")
    return <button className="CreateReview" onClick={onClick}>{buttonText}</button>;

  return <button className="buttonProfile" onClick={onClick}>{buttonText}</button>;
}

export default OpenModalButton;
