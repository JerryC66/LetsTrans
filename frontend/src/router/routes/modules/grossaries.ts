import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const DASHBOARD: AppRouteRecordRaw = {
  path: '/glossaries',
  name: 'glossariesHome',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.grossaries.management',
    requiresAuth: true,
    icon: 'icon-book',
    order: 0,
    showInMenu: true,
  },
  children: [
    {
      path: '',
      name: 'glossaryList',
      component: () => import('@/views/glossaries/index.vue'),
      meta: {
        locale: 'menu.grossaries.list',
        requiresAuth: true,
        roles: ['*'],
        showInMenu: true,
      },
    },
  ],
};

export default DASHBOARD;
