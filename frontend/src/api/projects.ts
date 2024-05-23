import request from './_request';
import { HttpRes } from '@/types/api';

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

export const deleteProject = (project_ids: number[]) => {
  return request({
    url: '/projects',
    method: 'DELETE',
    data: { ids: project_ids },
  }) as Promise<HttpRes<any>>;
};

export const getProjectDetail = (project_id: number) => {
  return request({
    url: '/projects/{project_id}',
    method: 'GET',
  }) as Promise<HttpRes<any>>;
};
