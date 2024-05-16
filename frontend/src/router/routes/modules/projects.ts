import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const DASHBOARD: AppRouteRecordRaw = {
  path: '/projects',
  name: 'projectsHome',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.projects.management',
    requiresAuth: true,
    icon: 'icon-apps',
    order: 0,
    showInMenu: true,
  },
  children: [
    {
      path: '',
      name: 'projectList',
      component: () => import('@/views/projects/projectsList/index.vue'),
      meta: {
        locale: 'menu.projects.list',
        requiresAuth: true,
        roles: ['*'],
        showInMenu: true,
      },
    },
    {
      path: ':projectId/files',
      name: 'manageFiles',
      component: () => import('@/views/projects/manageFiles/index.vue'),
      meta: {
        locale: 'menu.projects.manageFiles',
        requiresAuth: true,
        roles: ['*'],
        showInMenu: false,
        props: true,
      },
    },
    {
      path: ':projectId/files/:fileId/translate',
      name: 'translatePage',
      component: () => import('@/views/projects/translatePage/index.vue'),
      meta: {
        locale: 'menu.projects.translate',
        requiresAuth: true,
        roles: ['*'],
        showInMenu: false,
        props: true,
      },
    },
  ],
};

export default DASHBOARD;
