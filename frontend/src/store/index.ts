import { createPinia } from 'pinia';
import useAppStore from './modules/app';
import useUserStore from './modules/user';
import useTabBarStore from './modules/tab-bar';
import useTranslationStore from './modules/translation';

const pinia = createPinia();

export { useAppStore, useUserStore, useTabBarStore, useTranslationStore };
export default pinia;
