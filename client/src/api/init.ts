import axios, { type AxiosError, type InternalAxiosRequestConfig } from 'axios';
import { refresh } from './auth';
import { GET_COOKIE, SET_COOKIE } from '@/utils/functions';

// --- Константы ---
const API_URL = 'http://localhost:10801';
let isRefreshing = false;

// --- Расширение интерфейса Axios ---
interface IRetryConfig extends InternalAxiosRequestConfig {
  isRetry?: boolean;
}

// --- Экземпляр Axios ---
export const $api = axios.create({
  baseURL: API_URL,
  withCredentials: true,
});

// 1. Request Interceptor (для Access Token и логирования)
$api.interceptors.request.use((config: InternalAxiosRequestConfig) => {
  // Логирование
  console.log(`[API Request] -> Method: ${config.method?.toUpperCase()} | URL: ${config.baseURL}${config.url}`);

  // Прикрепляем Access Token
  const accessToken = GET_COOKIE('access_token');
  if (accessToken) {
    config.headers.Authorization = `Bearer ${accessToken}`;
  }

  return config;
}, (error) => {
  return Promise.reject(error);
});


// 2. Response Interceptor (для обработки 401 Unauthorized)
$api.interceptors.response.use((response) => {
  return response;
}, async (error: AxiosError) => {
  const originalRequest = error.config as IRetryConfig;

  // Проверяем ошибку 401 и что это не повторный запрос
  if (error.response?.status === 401 && originalRequest.url !== '/refresh' && !originalRequest.isRetry) {

    if (isRefreshing) {
      // Если токен уже обновляется, ждем и повторяем запрос
      return new Promise((resolve) => {
        // Используем setTimeout как простой пример очереди
        setTimeout(() => resolve($api(originalRequest)), 500);
      });
    }

    originalRequest.isRetry = true;
    isRefreshing = true;
    console.log('[API Interceptor] 401 Unauthorized. Attempting to refresh token...');

    try {
      // Запрос на /refresh. Ответ должен содержать новый AccessToken
      const response = await refresh();
      const newAccessToken = response.data.access_token;

      const expiryDate = new Date();
      expiryDate.setTime(expiryDate.getTime() + 5 * 60 * 1000);
      SET_COOKIE('access_token', newAccessToken, expiryDate);

      console.log('[API Interceptor] Token refreshed. Retrying original request.');
      // Повторяем оригинальный запрос
      return $api(originalRequest);
    } catch (e) {
      console.error('[API Interceptor] Refresh token failed. Logging out...');
      // В случае неудачи перенаправляем на логин
      window.location.href = '/auth'; 
      return Promise.reject(error);
    } finally {
      isRefreshing = false;
    }
  }

  return Promise.reject(error);
});