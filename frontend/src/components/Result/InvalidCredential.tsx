"use client";
import React from "react";
import { Button, Result } from "antd";
import { useRouter } from "next/navigation";

const InvalidCredential: React.FC = () => {
  const router = useRouter();

  const handleSubmit = () => {
    localStorage.removeItem("authToken");
    localStorage.removeItem("username");
    router.push("/");
  };

  return (
    <Result
      status="403"
      title="403"
      subTitle="Sorry, you are not authorized to access this page."
      extra={
        <Button type="primary" onClick={handleSubmit}>
          Back Home
        </Button>
      }
    />
  );
};

export default InvalidCredential;
