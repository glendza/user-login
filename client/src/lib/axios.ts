import Axios, { AxiosRequestConfig } from "axios";

import { API_URL } from "@/config";

const authRequestInterceptor = (config: AxiosRequestConfig) => {
  const token = localStorage.getItem('token')

  const headers = () => {
    if (!config.headers) {
      config.headers = {};
    }
    return config.headers as Record<string, string>
  }

  if (token) {
    headers().Authorization = token;
  }
  headers().Accept = "application/json";
  return config;
};

export const axios = Axios.create({
  baseURL: API_URL,
});

axios.interceptors.request.use(authRequestInterceptor);
