import { useState, FormEvent } from 'react';

import { useDispatch, useSelector } from 'react-redux';
import { Navigate } from 'react-router-dom';
import { login } from '../../Actions/userActions';
import { AppDispatch } from '../../store/store';
// Adjust path to the store if necessary

// Type for the errors state
interface Errors {
  credential?: string;
}

function LoginFormPage() {
  const dispatch = useDispatch<AppDispatch>();
  const sessionUser = useSelector((state: RootState) => state.session.user);
  const [credential, setCredential] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [errors, setErrors] = useState<Errors>({});

  if (sessionUser) return <Navigate to="/" replace={true} />;

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setErrors({});
    dispatch(login(credential,password))
  };

  return (
    <>
      <h1>Log In</h1>
      <form onSubmit={handleSubmit}>
        <label>
          Username or Email
          <input
            type="text"
            value={credential}
            onChange={(e) => setCredential(e.target.value)}
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
        {errors.credential && <p>{errors.credential}</p>}
        <button type="submit">Log In</button>
      </form>
    </>
  );
}

export default LoginFormPage;
