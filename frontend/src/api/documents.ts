import { HttpRes } from '@/types/api';
import request from './_request';
import { Document, Segment } from './model';

export const getDocumentSegments = (documentId: number) => {
  return request({
    url: `/documents/${documentId}/segments`,
    method: 'GET',
  }) as Promise<HttpRes<{ document: Document; segments: Segment[] }>>;
};
