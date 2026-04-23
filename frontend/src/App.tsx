import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/Home';
import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import ProtectedRoute from './components/ProtectedRoute'
import PublicRoute from './components/PublicRoute'

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<PublicRoute><Home /> </PublicRoute>} />
        <Route path="/login" element={<PublicRoute><Login /> </PublicRoute>} />
        <Route path="/register" element={<PublicRoute><Register /> </PublicRoute>} />
        <Route path="/dashboard" element={<ProtectedRoute><Dashboard /> </ProtectedRoute>} />
      </Routes>
    </Router>
  );
}

export default App;
