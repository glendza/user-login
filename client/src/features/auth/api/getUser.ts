import { axios } from '@/lib/axios';

import { UserProfileResponse } from '../types';

export const getUser = async (): Promise<UserProfileResponse> => {
  const { data } = await axios.get('/user/self');
  return data;
};
