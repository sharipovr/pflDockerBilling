// services/api.ts
import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://localhost:8080', // адрес твоего backend API
});

// Загружаем данные пользователя
export const getUser = async (userId: number) => {
  return apiClient.get(`/api/users/${userId}`);
};

// Загружаем подписки пользователя
export const getUserSubscriptions = async (userId: number) => {
  return apiClient.get(`/api/subscriptions?user_id=${userId}`);
};

// Загружаем данные пользователя по email
export const getUserByEmail = (email: string) => {
  return apiClient.get(`/api/users/email/${email}`);
};
