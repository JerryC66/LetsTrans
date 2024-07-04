import { HttpRes } from '@/types/api';
import request from './_request';

export const getDocumentSegments = (documentId: number) => {
  return request({
    url: `/documents/${documentId}/segments`,
    method: 'GET',
  }) as Promise<HttpRes<any>>;
};

export const updateDocumentSegment = (
  documentId: number,
  segmentId: number,
  data: {
    targetText: string;
    finished: boolean;
  },
) => {
  return request({
    url: `/documents/${documentId}/segments/${segmentId}`,
    method: 'PUT',
    data,
  }) as Promise<HttpRes<any>>;
};
