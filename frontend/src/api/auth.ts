import { CommonRes } from './common_res';
import axios from 'axios';

interface UserTokenData {
  token: string;
}

export function getUserToken() {
  return axios.get('/user/token');
}
