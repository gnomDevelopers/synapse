import { createRouter, createWebHistory, type RouteLocationNormalized } from 'vue-router'
import { useUserStore } from '@/stores/userStore';
import { getUser } from '@/api/api';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      name: "AuthPage",
      path: "/auth",
      meta: { requiresAuth: false },
      redirect: { name: 'AuthPageLogin' },
      component: () => import("@/pages/AuthPage/AuthPage.vue"),
      children: [
        {
          name: "AuthPageLogin",
          path: "login",
          meta: { requiresAuth: false },
          component: () => import("@/pages/AuthPage/widgets/loginForm.vue"),
        },
        {
          name: "AuthPageSignUp",
          path: "signup",
          meta: { requiresAuth: false },
          component: () => import("@/pages/AuthPage/widgets/signUpForm.vue"),
        },
      ]
    },
    {
      name: "MainPage",
      path: "/",
      meta: { requiresAuth: true },
      redirect: { name: 'DisciplinesPage' },
      component: () => import("@/pages/MainPage/MainPage.vue"),
      children: [
        {
          name: "DisciplinesPage",
          path: "disciplines",
          meta: { requiresAuth: true },
          component: () => import("@/pages/DisciplinesPage/DisciplinesPage.vue"),
        },
        {
          name: "DisciplineDetailsPage",
          path: "discipline/:id",
          meta: { requiresAuth: true },
          component: () => import("@/pages/DisciplineDetailsPage/DisciplineDetailsPage.vue"),
        },
        {
          name: "TestConstructorPage",
          path: "test/create",
          meta: { requiresAuth: true },
          component: () => import("@/pages/TestConstructorPage/TestConstructorPage.vue"),
        },
      ]
    },
  ],
});


// 2. Глобальный навигационный хук (Middleware)
router.beforeEach(async (to: RouteLocationNormalized, from: RouteLocationNormalized, next) => {
  // Получаем экземпляр стора
  const userStore = useUserStore(); 
  
  // Флаг, требующий авторизации для текущего маршрута
  const requiresAuth = to.meta.requiresAuth as boolean;

  // Проверка и загрузка данных пользователя (если 'isAuthorized' === null)
  if (userStore.isAuthorized === null) {
    try {
      const userData = await getUser(); 
      
      // Если запрос успешен и вернул данные, сохраняем их и устанавливаем авторизацию
      userStore.$patch({
        isAuthorized: true,
        id: userData.data.id,
        fio: userData.data.fio
      });

    } catch (error) {
      // Если запрос неудачен, считаем пользователя неавторизованным
      userStore.isAuthorized = false;
      console.log("Пользователь не авторизован или ошибка API:", error);
    }
  }

  // Обработка перехода в зависимости от статуса авторизации
  if (requiresAuth && !userStore.isAuthorized) {
    // 1. Маршрут требует авторизации И пользователь НЕ авторизован
    // Перенаправляем на страницу входа
    
    // next({ name: 'AuthPage' });
    next();
  } else if (userStore.isAuthorized && (to.name === 'Login')) {
    // 2. Пользователь АВТОРИЗОВАН и пытается перейти на страницу входа/регистрации
    // Перенаправляем его на домашнюю страницу или страницу профиля
    next({ name: 'MainPage' }); 
  } else {
    // 3. Во всех остальных случаях (авторизован на защищенную, или неавторизован на публичную)
    // Разрешаем переход
    next();
  }
});


export default router
