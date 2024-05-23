// src/api/auth.ts
import request from './_request';
import { HttpRes } from '@/types/api';

export const getToken = () => {
  return request({
    url: '/user/token',
    method: 'GET',
  }) as Promise<HttpRes<{ token: string }>>;
};
