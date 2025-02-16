// src/features/auth/authService.ts
import axios from "axios";
import { AuthResponse } from "../types/auth";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export const loginUser = async (
  username: string,
  password: string
): Promise<AuthResponse> => {
  try {
    const response = await axios.post<AuthResponse>(`${API_URL}/login`, {
      username,
      password,
    });
    const { token } = response.data;

    localStorage.setItem("authToken", token);

    return response.data;
  } catch (error) {
    throw new Error("Login failed! Please check your credentials.");
  }
};
