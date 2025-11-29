import {
  type IAuthRequest, 
  type IAuthResponse, 
  type IUser, 
  login, 
  logout, 
  refresh, 
  register, 
  getUser 
} from './auth'; // Импортируем все из auth.ts

// --- Экспорт DTO и методов ---

// Экспортируем все импортированные функции и интерфейсы из auth.ts
export type { IAuthRequest, IAuthResponse, IUser };
export { login, logout, refresh, register, getUser };