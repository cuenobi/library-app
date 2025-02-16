import React from "react";
import { Alert } from "antd";

interface AlertProps {
  message: string;
}

const ErrorAlert: React.FC<AlertProps> = ({ message }) => (
  <>
    <Alert message={message} type="error" />
    <br />
  </>
);

export default ErrorAlert;
