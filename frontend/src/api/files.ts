import { HttpRes } from '@/types/api';
import request from './_request';

interface FileUploadResponse {
  id: number;
  filename: string;
  filetype: string;
}

// export const uploadFile = (fileData: FormData) => {
//   return request({
//     url: '/files',
//     method: 'POST',
//     data: fileData,
//     headers: { 'Content-Type': 'multipart/form-data' },
//   }) as Promise<HttpRes<FileUploadResponse>>;
// };

export const addFileToProject = (projectId: number, data: FormData) => {
  return request({
    url: `/projects/${projectId}/files`,
    method: 'POST',
    data,
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  }) as Promise<HttpRes<any>>;
};

export const deleteFileFromProject = (projectId: number, fileIds: number[]) => {
  return request({
    url: `/projects/${projectId}/files`,
    method: 'DELETE',
    data: { ids: fileIds },
  }) as Promise<HttpRes<any>>;
};

export const downloadFiles = (
  projectId: number,
  documentIds: number[],
  type: 'origin' | 'translated',
) => {
  return request({
    url: `/projects/${projectId}/files`,
    method: 'GET',
    data: { ids: documentIds, type },
  }) as Promise<HttpRes<any>>;
};
