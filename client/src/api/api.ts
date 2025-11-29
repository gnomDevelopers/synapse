import {
  type IAuthRequest, 
  type IAuthResponse, 
  type IUserResponse, 
  login, 
  logout, 
  refresh, 
  register, 
  getUser 
} from './auth'; // Импортируем все из auth.ts

// --- Экспорт DTO и методов ---

// Экспортируем все импортированные функции и интерфейсы из auth.ts
export type { IAuthRequest, IAuthResponse, IUserResponse };
export { login, logout, refresh, register, getUser };