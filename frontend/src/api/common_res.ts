export interface CommonRes<T> {
  code: number;
  msg: string;
  data: T;
}
