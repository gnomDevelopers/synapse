import { type AxiosResponse } from 'axios';
import { $api } from './init';

// --- Интерфейсы (DTO) ---

/**
 * Интерфейс пользователя (для ответа и для получения данных)
 */
export interface IUser {
  id: number;
  first_name: string;
  last_name: string;
  third_name?: string; // Отчество может быть необязательным
  email: string;
}

/**
 * Интерфейс тела запроса для регистрации/логина
 */
export interface IAuthRequest {
  first_name?: string;
  last_name?: string;
  third_name?: string;
  email: string;
  password: string;
}

/**
 * Интерфейс ответа при успешной авторизации/обновлении
 */
export interface IAuthResponse {
  user: IUser;
  accessToken: string;
}

// --- Методы API ---

/**
 * POST /register - Регистрация пользователя
 * @param data Данные пользователя
 */
export const register = (data: IAuthRequest): Promise<AxiosResponse<IAuthResponse>> => {
  return $api.post<IAuthResponse>('/auth/register', data);
};

/**
 * POST /login - Вход пользователя
 * @param data Email и пароль
 */
export const login = (data: Pick<IAuthRequest, 'email' | 'password'>): Promise<AxiosResponse<IAuthResponse>> => {
  return $api.post<IAuthResponse>('/auth/login', data);
};

/**
 * GET /user - Получение данных текущего пользователя (требует Access Token)
 */
export const getUser = (): Promise<AxiosResponse<IUser>> => {
  return $api.get<IUser>('/auth/user');
};

/**
 * GET /refresh - Обновление токенов (используется интерцептором)
 */
export const refresh = (): Promise<AxiosResponse<IAuthResponse>> => {
  return $api.get<IAuthResponse>('/auth/refresh');
};

/**
 * POST /logout - Выход пользователя
 */
export const logout = (): Promise<AxiosResponse> => {
  return $api.post('/auth/logout');
};
