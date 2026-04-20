import { useEffect } from 'react';

export default function Dashboard() {
  const token = localStorage.getItem('token');
  const username: string = token ? JSON.parse(atob(token.split('.')[1])).username : '';

  useEffect(() => {
    if (!token) {
      window.location.href = '/login';
    }
  }, [token]);

  if (!token) return null;

  return (
    <div style={{ textAlign: 'center', padding: '2rem' }}>
      <h1>Welcome, {username}!</h1>
    </div>
  );
}
