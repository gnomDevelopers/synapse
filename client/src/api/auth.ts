import { type AxiosResponse } from 'axios';
import { $api } from './init';

// --- Интерфейсы (DTO) ---

/**
 * Интерфейс пользователя (для ответа и для получения данных)
 */
export interface IUserResponse {
  id: number;
  fio: string;
  email: string;
}

/**
 * Интерфейс тела запроса для регистрации/логина
 */
export interface IAuthRequest {
  full_name?:string;
  email: string;
  password: string;
}

/**
 * Интерфейс ответа при успешной авторизации/обновлении
 */
export interface IAuthResponse {
  access_token: string;
  refresh_token: string;
}

// --- Методы API ---

/**
 * POST /register - Регистрация пользователя
 * @param data Данные пользователя
 */
export const register = (data: IAuthRequest): Promise<AxiosResponse<IAuthResponse>> => {
  return $api.post<IAuthResponse>('/sign-up', data);
};

/**
 * POST /login - Вход пользователя
 * @param data Email и пароль
 */
export const login = (data: Pick<IAuthRequest, 'email' | 'password'>): Promise<AxiosResponse<IAuthResponse>> => {
  return $api.post<IAuthResponse>('/login', data);
};

/**
 * GET /user - Получение данных текущего пользователя (требует Access Token)
 */
export const getUser = (): Promise<AxiosResponse<IUserResponse>> => {
  return $api.get<IUserResponse>('/auth/user');
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
