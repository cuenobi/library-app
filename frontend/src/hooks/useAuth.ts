import { useState } from "react";
import { loginUser } from "../features/authService";
import { AuthResponse } from "../types/auth";
import { message } from "antd";

export const useAuth = () => {
  const [authToken, setAuthToken] = useState<string | null>(null);

  const login = async (username: string, password: string) => {
    try {
      const response: AuthResponse = await loginUser(username, password);
      if (response.token) {
        setAuthToken(response.token);
        return {
          success: true,
          role: response.role,
          message: response.message,
        };
      }
      return { success: false, role: null, message: response.message };
    } catch (error) {
      return { success: false, role: null, message: error };
    }
  };

  return { authToken, login, message };
};
