import { HttpRes } from '@/types/api';
import request from './_request';

export const getProjects = () => {
  return request({
    url: '/projects',
    method: 'GET',
  }) as Promise<HttpRes<any[]>>;
};

export const createProject = (data: {
  source_lang: string;
  target_lang: string;
  name: string;
  deadline: string;
  comment: string;
}) => {
  return request({
    url: '/projects',
    method: 'POST',
    data,
  }) as Promise<HttpRes<any>>;
};

export const deleteProject = (projectIds: number[]) => {
  return request({
    url: '/projects',
    method: 'DELETE',
    data: { ids: projectIds },
  }) as Promise<HttpRes<any>>;
};

export const getProjectDetail = (projectId: number) => {
  return request({
    url: `/projects/${projectId}`,
    method: 'GET',
  }) as Promise<HttpRes<any>>;
};
