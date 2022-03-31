import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/flow',
      name: 'flow',
      component: () => import('../views/FlowView.vue'),
    },
    {
      path: '/code',
      name: 'code',
      component: () => import('../views/CodeView.vue'),
    },
  ],
});

export default router;
