import request from './_request';
import { HttpRes } from '@/types/api';

interface FileUploadResponse {
  id: number;
  filename: string;
  filetype: string;
}

export const uploadFile = (fileData: FormData) => {
  return request({
    url: '/files',
    method: 'POST',
    data: fileData,
    headers: { 'Content-Type': 'multipart/form-data' },
  }) as Promise<HttpRes<FileUploadResponse>>;
};

export const addFileToProject = (project_id: number, file_ids: number[]) => {
  return request({
    url: '/projects/{project_id}/files',
    method: 'POST',
    data: { ids: file_ids },
  }) as Promise<HttpRes<any>>;
};

export const deleteFileFromProject = (
  project_id: number,
  file_ids: number[]
) => {
  return request({
    url: '/projects/{project_id}/files',
    method: 'DELETE',
    data: { ids: file_ids },
  }) as Promise<HttpRes<any>>;
};

export const downloadFiles = (
  project_id: number,
  document_ids: number[],
  type: 'origin' | 'translated'
) => {
  return request({
    url: `/projects/${project_id}/files`,
    method: 'GET',
    data: { ids: document_ids, type },
  }) as Promise<HttpRes<any>>;
};
