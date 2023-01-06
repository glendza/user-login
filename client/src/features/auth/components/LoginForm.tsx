import React from "react";
import { useForm } from "react-hook-form";
import { useLogin } from "@/lib/auth";
import { LoginCredentials } from "../api/login";

export const LoginForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginCredentials>();

  const login = useLogin();

  const onSubmit = (data: LoginCredentials) => {
    login.mutate(data);
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <input {...register("username", { required: true })} />
      {errors.username && <span>This field is required</span>}
      <input {...register("password", { required: true })} />
      {errors.password && <span>This field is required</span>}
      <input type="submit" value="Login" />
      {login.isError && <span>Invalid credentials</span>}
    </form>
  );
};
