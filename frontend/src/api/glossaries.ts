import { HttpRes } from '@/types/api';
import request from './_request';

export const getGlossaries = () => {
  return request({
    url: '/glossaries',
    method: 'GET',
  }) as Promise<HttpRes<any[]>>;
};

export const createGlossary = (data: {
  comment: string;
  author: string;
  name: string;
  source_lang: string;
  target_lang: string;
}) => {
  return request({
    url: '/glossaries',
    method: 'POST',
    data,
  }) as Promise<HttpRes<any>>;
};

export const deleteGlossary = (glossaryId: string) => {
  return request({
    url: `/glossaries/${glossaryId}`,
    method: 'DELETE',
  }) as Promise<HttpRes<any>>;
};

export const getGlossaryTerms = (glossaryId: number) => {
  return request({
    url: `/glossaries/${glossaryId}`,
    method: 'GET',
  }) as Promise<HttpRes<any>>;
};

export const addTermToGlossary = (
  glossaryId: number,
  data: {
    source_lang: string;
    target_lang: string;
    source_text: string;
    target_text: string;
  },
) => {
  return request({
    url: `/glossaries/${glossaryId}/terms`,
    method: 'POST',
    data,
  }) as Promise<HttpRes<any[]>>;
};

export const importGossaryCSV = (glossaryId: string, file: FormData) => {
  return request({
    url: `/glossaries/${glossaryId}/terms/batch`,
    method: 'POST',
    data: file,
  }) as Promise<HttpRes<any[]>>;
};

export const updateGlossaryTerm = (
  termId: number,
  data: {
    source_lang: string;
    target_lang: string;
    source_text: string;
    target_text: string;
  },
) => {
  return request({
    url: `/glossaries/terms/${termId}`,
    method: 'PUT',
  }) as Promise<HttpRes<any[]>>;
};

export const getGlossarySuggestion = (sourceText: string) => {
  return request({
    url: `/glossaries/suggestion?source_text=${sourceText}`,
    method: 'GET',
  }) as Promise<HttpRes<any>>;
};
