import { Link } from 'react-router-dom';

export default function Home() {
  return (
    <div style={{ textAlign: 'center', padding: '2rem' }}>
      <h1>Welcome to Trivia Game</h1>
      <p>Test your knowledge with friends!</p>
      <Link to="/register">
        <button>Register</button>
      </Link>
      <Link to="/login" style={{ marginLeft: '1rem' }}>
        <button>Login</button>
      </Link>
    </div>
  );
}