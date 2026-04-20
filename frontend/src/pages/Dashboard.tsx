import { useEffect, useState } from 'react';

export default function Dashboard() {
  const [username, setUsername] = useState<string>('');

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (!token) {
      window.location.href = '/login';
      return;
    }

    const payload = JSON.parse(atob(token.split('.')[1]));
    setUsername(payload.username);
  }, []);

  return (
    <div style={{ textAlign: 'center', padding: '2rem' }}>
      <h1>Welcome, {username}!</h1>
    </div>
  );
}
