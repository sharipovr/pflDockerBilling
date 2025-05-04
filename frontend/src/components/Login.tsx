import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Login: React.FC = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleLogin = (e: React.FormEvent) => {
    e.preventDefault();
    // Тут заглушка авторизации без реальной проверки
    if (username && password) {
      // Сохраняем email (username) в localStorage перед редиректом
      localStorage.setItem('userEmail', username);
      navigate('/dashboard');
    }
  };

  return (
    <div className="h-screen flex items-center justify-center">
      <form className="bg-white p-6 rounded shadow-md" onSubmit={handleLogin}>
        <h2 className="mb-4 text-xl font-bold">Login</h2>
        <input
          className="mb-2 w-full border p-2"
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <input
          className="mb-4 w-full border p-2"
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button className="w-full bg-blue-500 text-white p-2 rounded">
          Войти
        </button>
      </form>
    </div>
  );
};

export default Login;