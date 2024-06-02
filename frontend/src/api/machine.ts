import { HttpRes } from '@/types/api';
import request from './_request';

export const getTransAPI = () => {
  return request({
    url: '/machine/engines',
    method: 'GET',
  }) as Promise<HttpRes<any>>;
};

export const postTransAPI = (data: { engine: string; document_id: number }) => {
  return request({
    url: '/machine/translate',
    method: 'POST',
    data,
  }) as Promise<HttpRes<any>>;
};
