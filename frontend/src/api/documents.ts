import { CommonRes } from './common_res';
import { Document, Segment } from './model';
import axios from 'axios';

export function getDocumentSegments(
  document_id: string,
  segments: Segment[]
): Promise<CommonRes<{}>> {
  return axios.get('/documents/{document_id}/segments');
}

export function updateDocumentSegment(
  document_id: string,
  segment_id: string,
  target_text: string,
  finished: boolean
): Promise<CommonRes<{}>> {
  return axios.post('/documents/{document_id}/segments/{segment_id}');
}
