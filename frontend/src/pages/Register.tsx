import { useState } from 'react';

export default function Register() {
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [username, setUsername] = useState<string>('');

  function handleRegister(e: SubmitEvent) {
    e.preventDefault();
    console.log('Registering:', { email, password, username });

    fetch('http://localhost:8080/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password, username }),
    })
      .then(res => {
        if (!res.ok) {
          throw new Error('Registration failed');
        }
        return res.json();
      })
      .then(data => {
        localStorage.setItem('token', data.token);
        window.location.href = '/dashboard';
      }).catch(err => {
        alert(err.message);
      });
  }

  return (
    <div style={{ textAlign: 'center', padding: '2rem' }}>
      <h2>Register</h2>
      <form onSubmit={handleRegister}>
      <input
          type="text"
          placeholder="username"
          value={username}
          required
          onChange={(e) => setUsername(e.target.value)}
        /><br /><br />
        <input
          type="email"
          placeholder="Email"
          value={email}
          required
          onChange={(e) => setEmail(e.target.value)}
        /><br /><br />
        <input
          type="password"
          placeholder="Password"
          value={password}
          required
          onChange={(e) => setPassword(e.target.value)}
        /><br /><br />
        <button type="submit">Register</button>
      </form>
    </div>
  );
}
