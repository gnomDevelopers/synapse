import { createRouter, createWebHistory } from 'vue-router'

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
})

export default router
