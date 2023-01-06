import { configureAuth } from "react-query-auth";

import {
  getUser,
  login,
  LoginCredentials,
  UserProfileResponse,
} from "@/features/auth";

const userFn = async (): Promise<UserProfileResponse | null> => {
  if (localStorage.getItem("token") !== null) {
    const data = await getUser();
    return data;
  }
  return null;
};

const loginFn = async (
  credentials: LoginCredentials
): Promise<UserProfileResponse | null> => {
  const { token } = await login(credentials);
  localStorage.setItem("token", token);
  const user = await userFn();
  return user;
};

const registerFn = async (credentials: unknown) => {
  throw new Error("Not implemented!");
};

async function logoutFn() {
  localStorage.removeItem("token");
  window.location.assign((window.location.origin as unknown) as string);
}

const authConfig = {
  userFn,
  loginFn,
  registerFn,
  logoutFn,
};

export const { AuthLoader, useUser, useLogin } = configureAuth<
  UserProfileResponse | null,
  unknown,
  LoginCredentials,
  unknown
>(authConfig);
