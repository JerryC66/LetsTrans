import { HttpRes } from '@/types/api';
import request from './_request';

export const getMemorySuggestion = (sourceText: string) => {
  return request({
    url: `/memories/suggestion?source_text=${sourceText}`,
    method: 'GET',
  }) as Promise<HttpRes<any>>;
};
