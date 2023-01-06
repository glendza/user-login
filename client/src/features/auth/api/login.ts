import { axios,  } from "@/lib/axios";
import { AuthResponse } from "../types";


export type LoginCredentials = {
  username: string;
  password: string;
}


export const login = async (credentials: LoginCredentials): Promise<AuthResponse> => {
  const { data } = await axios.post("/auth/login", credentials);
  return data;
};
