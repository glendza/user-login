import { axios } from "@/lib/axios";
import { AuthResponse } from "../types";


export const login = (username: string, password: string): Promise<AuthResponse> => {
  const data = { username, password };
  return axios.post("/auth/login", data);
};
