// components/Dashboard.tsx
import React, { useEffect, useState } from 'react';
import Navbar from './Navbar';
import { getUserByEmail, getUserSubscriptions } from '../services/api';

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
        const userEmail = localStorage.getItem('userEmail');
        if (!userEmail) {
          setError('Пользователь не авторизован.');
          setLoading(false);
          return;
        }

        // Получаем данные пользователя по email
        const userResp = await getUserByEmail(userEmail);
        setUser(userResp.data);

        // Теперь загружаем подписки по полученному ID пользователя
        const subsResp = await getUserSubscriptions(userResp.data.ID);
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