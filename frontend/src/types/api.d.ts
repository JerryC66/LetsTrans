export interface HttpRes<T> {
  code: number;
  msg: string;
  data: T;
}
