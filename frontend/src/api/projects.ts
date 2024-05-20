import { CommonRes } from './common_res';
import { Project, Document } from './model';
import axios from 'axios';

export function createProject(
  projectData: Project
): Promise<CommonRes<Project>> {
  return axios.post('/projects');
}

export function getProjects(): Promise<CommonRes<Project[]>> {
  return axios.get('/projects');
}

export function deleteProjects(project_ids: string[]): Promise<CommonRes<{}>> {
  return axios.delete('/projects');
}
