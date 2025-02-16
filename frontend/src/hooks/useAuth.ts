import { useState } from "react";
import { loginUser } from "../features/authService";
import { AuthResponse } from "../types/auth";

export const useAuth = () => {
  const [authToken, setAuthToken] = useState<string | null>(null);

  const login = async (username: string, password: string) => {
    try {
      const response: AuthResponse = await loginUser(username, password);
      if (response.token) {
        setAuthToken(response.token);
        return true;
      }
      return false;
    } catch (error) {
      console.error(error);
      return false;
    }
  };

  return { authToken, login };
};
