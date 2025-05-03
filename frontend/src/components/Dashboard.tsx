// components/Dashboard.tsx
import React, { useEffect, useState } from 'react';
import Navbar from './Navbar';
import { getUser, getUserSubscriptions } from '../services/api';

interface User {
  id: number;
  name: string;
  email: string;
}

interface Subscription {
  id: number;
  plan: string;
  status: string;
}

const Dashboard: React.FC = () => {
  const [user, setUser] = useState<User | null>(null);
  const [subscriptions, setSubscriptions] = useState<Subscription[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>('');

  useEffect(() => {
    const loadData = async () => {
      try {
        const userId = 1; // предполагаемый id текущего пользователя
        const [userResp, subsResp] = await Promise.all([
          getUser(userId),
          getUserSubscriptions(userId),
        ]);

        setUser(userResp.data);
        setSubscriptions(subsResp.data);
      } catch (err) {
        console.error(err);
        setError('Ошибка загрузки данных.');
      } finally {
        setLoading(false);
      }
    };

    loadData();
  }, []);

  if (loading) return <div className="p-6">Загрузка...</div>;
  if (error) return <div className="p-6 text-red-500">{error}</div>;

  return (
    <div>
      <Navbar />
      <div className="container mx-auto p-6">
        <h1 className="text-2xl font-bold mb-4">Dashboard</h1>
        <div className="mb-4">
          <h2 className="font-semibold">Информация о пользователе:</h2>
          <p>Имя: {user?.name}</p>
          <p>Email: {user?.email}</p>
        </div>
        <div>
          <h2 className="font-semibold">Подписки:</h2>
          <ul>
            {subscriptions.map((sub, index) => (
              <li key={sub.id || index}>
                {sub.plan} - <strong>{sub.status}</strong>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;