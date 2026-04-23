import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom'

export default function Dashboard() {
  const navigate = useNavigate();
  const token = localStorage.getItem('token');
  const username: string = token ? JSON.parse(atob(token.split('.')[1])).username : '';

  function handleLogout() {                                                                                                                                                                                          
    localStorage.removeItem('token');                                                                                                                                                                                
    navigate('/');                                                                                                                                                                                      
  }

  useEffect(() => {
    if (!token) {
      navigate('/login');
    }
  }, [token]);

  if (!token) return null;

  return (
    <div style={{ textAlign: 'center', padding: '2rem' }}>
      <h1>Welcome, {username}!</h1>
      <button onClick={handleLogout}>Logout</button>
    </div>
  );
}
