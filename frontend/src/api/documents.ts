import request from './_request';
import { HttpRes } from '@/types/api';
import { Document, Segment } from './model';

export const getDocumentSegments = (documentId: string) => {
  return request({
    url: `/documents/${documentId}/segments`,
    method: 'GET',
  }) as Promise<HttpRes<{ document: Document; segments: Segment[] }>>;
};
