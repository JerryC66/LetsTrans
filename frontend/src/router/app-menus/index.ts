import { appRoutes, appExternalRoutes } from '../routes';

const filterAndMapRoutes = (routes: any[]): any[] => {
  return routes
    .filter((route) => route.meta?.showInMenu)
    .map((el) => {
      const { name, path, meta, redirect, children } = el;
      // 递归处理子路由
      const processedChildren = children
        ? filterAndMapRoutes(children)
        : undefined;
      return {
        name,
        path,
        meta,
        redirect,
        children: processedChildren,
      };
    });
};

const mixinRoutes = [
  ...filterAndMapRoutes(appRoutes),
  ...filterAndMapRoutes(appExternalRoutes),
];

const appClientMenus = mixinRoutes.map((el) => {
  const { name, path, meta, redirect, children } = el;
  return {
    name,
    path,
    meta,
    redirect,
    children,
  };
});

export default appClientMenus;
