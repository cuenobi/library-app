import React from "react";
import { Button, message } from "antd";
import axios from "axios";
import { SwapLeftOutlined } from "@ant-design/icons";

interface ReturnButtonProps {
  userID: string;
  bookID: string;
  onReturnSuccess: () => void;
}

const ReturnButton: React.FC<ReturnButtonProps> = ({
  userID,
  bookID,
  onReturnSuccess,
}) => {
  const handleReturn = async () => {
    try {
      const token = localStorage.getItem("authToken");
      if (!token) {
        message.error("No authorization token found.");
        return;
      }

      await axios.post(
        `http://localhost:8080/book/return/${userID}/${bookID}`,
        {},
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      message.success("Book returned successfully!");

      onReturnSuccess();
    } catch (error) {
      message.error("Failed to return book.");
    }
  };

  return (
    <Button
      type="default"
      shape="round"
      icon={<SwapLeftOutlined />}
      onClick={handleReturn}
    >
      Return
    </Button>
  );
};

export default ReturnButton;
