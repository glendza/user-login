import React from 'react';
import { useRoutes } from "react-router-dom";

import { Home } from "@/features/home";

export const AppRoutes = () => {
  const routes = [{ path: "/", element: <Home /> }];
  const element = useRoutes(routes);
  return <>{element}</>;
};
