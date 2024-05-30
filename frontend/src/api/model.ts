export interface Project {
  source_lang: string;
  target_lang: string;
  name: string;
  deadline: string;
  comment: string;
}

export interface Document {
  id: number;
  created_at: string;
  updated_at: string;
  name: string;
  progress: number;
  filetype: string;
  author: string;
  projrct_id: number;
}

export interface Segment {
  document_id: string;
  id: number;
  created_at: string;
  updated_at: string;
  source_text: string;
  target_text: string;
  finished: boolean;
}
