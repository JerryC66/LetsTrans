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
};

export default DASHBOARD;
