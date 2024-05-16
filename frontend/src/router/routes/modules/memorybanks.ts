import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const DASHBOARD: AppRouteRecordRaw = {
  path: '/memorybanks',
  name: 'memorybanksHome',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.memorybanks.management',
    requiresAuth: true,
    icon: 'icon-common',
    order: 0,
    showInMenu: true,
  },
};

export default DASHBOARD;
