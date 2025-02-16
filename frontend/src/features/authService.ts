import axios from "axios";
import { AuthResponse } from "../types/auth";
import { useRouter } from "next/navigation";

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

    const { token, role } = response.data;

    // Save token and role to localStorage
    localStorage.setItem("authToken", token);
    localStorage.setItem("role", role);

    return response.data;
  } catch (error: any) {
    // Check if error has response data
    if (error.response && error.response.data) {
      // If error has message in response, return that message
      throw new Error(error.response.data.message || "Login failed! Please check your credentials.");
    }

    // If there is no response from server, use default error message
    throw new Error("Login failed! Please check your credentials.");
  }
};