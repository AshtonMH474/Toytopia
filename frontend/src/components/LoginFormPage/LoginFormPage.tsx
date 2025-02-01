import { useState, FormEvent } from 'react';

import { useDispatch } from 'react-redux';
// import { Navigate } from 'react-router-dom';

import { AppDispatch } from '../../store/store';
import { useModal } from '../../Context/Modal';
import { login } from '../../store/session';
// Adjust path to the store if necessary

// Type for the errors state
interface Errors {
  email?: string;
}

function LoginFormPage() {
const {closeModal} = useModal();
  const dispatch = useDispatch<AppDispatch>();

  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [errors, setErrors] = useState<Errors>({});

//   if (sessionUser) return <Navigate to="/" replace={true} />;

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setErrors({});
    await dispatch(login({email,password}))
    await closeModal()
  };

  return (
    <>
      <h1>Log In</h1>
      <form onSubmit={handleSubmit}>
        <label>
          Username or Email
          <input
            type="text"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </label>
        <label>
          Password
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </label>
        {errors.email && <p>{errors.email}</p>}
        <button type="submit">Log In</button>
      </form>
    </>
  );
}

export default LoginFormPage;
