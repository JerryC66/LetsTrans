import request from './_request';
import { HttpRes } from '@/types/api';

export const getGlossaries = () => {
  return request({
    url: '/glossaries',
    method: 'GET',
  }) as Promise<HttpRes<any[]>>;
};

export const createGlossary = (data: {
  comment: string;
  author: string;
  name: string;
}) => {
  return request({
    url: '/glossaries',
    method: 'POST',
    data,
  }) as Promise<HttpRes<any>>;
};
