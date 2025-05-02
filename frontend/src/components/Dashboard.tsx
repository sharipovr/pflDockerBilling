import React from 'react';
import Navbar from './Navbar';

const Dashboard: React.FC = () => {
  const user = { username: 'test_user', email: 'user@example.com' };
  const subscriptions = [
    { id: 1, name: 'Docker Basic', status: 'active' },
    { id: 2, name: 'Docker Pro', status: 'expired' },
  ];

  return (
    <div>
      <Navbar />
      <div className="container mx-auto p-6">
        <h1 className="text-2xl font-bold mb-4">Dashboard</h1>
        <div className="mb-4">
          <h2 className="font-semibold">Информация о пользователе:</h2>
          <p>Имя: {user.username}</p>
          <p>Email: {user.email}</p>
        </div>

        <div>
          <h2 className="font-semibold">Подписки:</h2>
          <ul>
            {subscriptions.map((sub) => (
              <li key={sub.id}>
                {sub.name} - <strong>{sub.status}</strong>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;