import axios from "axios";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export const createBorrow = async (
  userID: string,
  bookID: string
): Promise<any> => {
  try {
    const response = await axios.post(`${API_URL}/borrow/${userID}/${bookID}`, {
      headers: {
        Authorization: localStorage.getItem("authToken"),
      },
    });
    return response.data;
  } catch (error: any) {
    throw new Error(error.response.data.message);
  }
};
