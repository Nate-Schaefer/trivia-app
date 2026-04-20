export default function Dashboard() {
  const token = localStorage.getItem('token');

  if (!token) {
    window.location.href = '/login';
    return null;
  }

  const payload = JSON.parse(atob(token.split('.')[1]));
  const username: string = payload.username;

  return (
    <div style={{ textAlign: 'center', padding: '2rem' }}>
      <h1>Welcome, {username}!</h1>
    </div>
  );
}
