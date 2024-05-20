import { CommonRes } from './common_res';
import axios from 'axios';

interface FileUploadData {
  file_id: string;
  filename: string;
  filetype: string;
}
export function uploadFile(file: File): Promise<CommonRes<FileUploadData>> {
  return axios.post('/files');
}

export function addFileToProject(
  project_id: string,
  file_id: string
): Promise<CommonRes<{}>> {
  return axios.post('/projects/{project_id}/files');
}

export function deleteFileFromProject(
  project_id: string,
  file_ids: string[]
): Promise<CommonRes<{}>> {
  return axios.delete('/projects/{project_id}/files');
}
