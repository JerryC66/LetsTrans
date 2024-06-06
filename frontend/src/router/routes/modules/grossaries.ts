import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const DASHBOARD: AppRouteRecordRaw = {
  path: '/glossaries',
  name: 'glossariesHome',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.glossaries.management',
    requiresAuth: true,
    icon: 'icon-book',
    order: 0,
    showInMenu: true,
  },
  children: [
    {
      path: '',
      name: 'glossaryList',
      component: () => import('@/views/glossaries/glossariesList/index.vue'),
      meta: {
        locale: 'menu.glossaries.list',
        requiresAuth: true,
        roles: ['*'],
        showInMenu: true,
      },
    },
    {
      path: ':glossaryId/terms',
      name: 'glossaryTerms',
      component: () => import('@/views/glossaries/glossaryTerms/index.vue'),
      meta: {
        locale: 'menu.glossaries.glossaryTerms',
        requiresAuth: true,
        roles: ['*'],
        showInMenu: false,
        props: true,
      },
    },
  ],
};

export default DASHBOARD;
