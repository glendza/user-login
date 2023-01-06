import * as React from "react";
import { QueryClientProvider } from '@tanstack/react-query';
import { BrowserRouter as Router } from "react-router-dom";
import { AuthLoader } from "@/lib/auth";
import { Login } from "@/features/auth"

import { queryClient } from '@/lib/react-query';

type AppProviderProps = {
  children: React.ReactNode;
};

export const AppProvider = ({ children }: AppProviderProps) => {
  return (
    <QueryClientProvider client={queryClient}>
      <AuthLoader
        renderLoading={() => <div>Loading ...</div>}
        renderUnauthenticated={() => <Login />}
      >
        <Router>{children}</Router>
      </AuthLoader>
    </QueryClientProvider>
  );
};
