import axios, { AxiosRequestConfig, AxiosInstance } from 'axios';

import { HttpRes } from '@/types/api.d';
import { Message } from '@arco-design/web-vue';
import i18n from '@/locale';
import { BASE_URL } from '@/constants';

// const { t } = useI18n();

export default function request<T = object>(config: AxiosRequestConfig) {
  const instance: AxiosInstance = axios.create({
    baseURL: BASE_URL,
    timeout: 60000,
    withCredentials: true,
  });

  instance.interceptors.request.use(
    (internalConfig: AxiosRequestConfig): AxiosRequestConfig => {
      return internalConfig;
    },
    (err: any): any => {
      console.error(err);
      Message.error(err.message);
    }
  );

  instance.interceptors.response.use(
    (response) => {
      const data = response.data || {};
      if (response.status === 200) {
        return data;
      }
      if (data.msg) {
        Message.error(data.msg);
      } else {
        console.error('# error', { response });
        Message.error(i18n.global.t('request.unknowErr'));
      }
      return null;
    },
    (err: any): any => {
      console.error(err);
      if (err.response?.data?.msg) {
        Message.error(err.response.data.msg);
      } else {
        Message.error(err.message);
      }
      return null;
    }
  );

  return new Promise<HttpRes<T>>((resolve, reject) => {
    instance(config).then(
      (res) => {
        resolve(res as unknown as HttpRes<T>);
      },
      (reason) => {
        reject(reason);
      }
    );
  });
}
