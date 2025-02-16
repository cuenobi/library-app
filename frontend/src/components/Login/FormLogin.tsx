"use client";
import { LockOutlined, UserOutlined } from "@ant-design/icons";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { LoginForm } from "@ant-design/pro-components";
import { Button, Form, Input, Checkbox, message } from "antd";
import { useAuth } from "../../hooks/useAuth";

export default () => {
  const { login } = useAuth();
  const router = useRouter();

  const handleRegister = (
    e: React.MouseEvent<HTMLAnchorElement, MouseEvent>
  ) => {
    e.preventDefault();
    router.push("/register"); // ใช้ router.push แทน href
  };

  const handleSubmit = async (values: {
    username: string;
    password: string;
  }) => {
    const { username, password } = values;

    const { success, role } = await login(username, password);

    if (success) {
      message.success("Login successful!");
      window.dispatchEvent(new Event("storage"));
      localStorage.setItem("username", username);
      if (role) {
        localStorage.setItem("role", role);
      }

      if (role === "2") {
        router.push("/dashboard");
      } else {
        router.push("/home");
      }
    } else {
      message.error("Login failed! Please check your credentials.");
    }
  };

  return (
    <LoginForm
      logo="./group.png"
      title="Signin"
      subTitle=" "
      name="signin"
      onFinish={handleSubmit}
      layout="vertical"
      initialValues={{ remember: true }}
      submitter={false}
    >
      <Form.Item
        name="username"
        label="Username"
        rules={[{ required: true, message: "Please enter your username!" }]}
      >
        <Input prefix={<UserOutlined />} placeholder="Enter your username" />
      </Form.Item>
      <Form.Item
        name="password"
        label="Password"
        rules={[{ required: true, message: "Please enter your password!" }]}
      >
        <Input.Password
          prefix={<LockOutlined />}
          placeholder="Enter your password"
        />
      </Form.Item>
      <Form.Item name="remember" valuePropName="checked">
        <div>
          <Checkbox>Remember me</Checkbox>
          <a
            style={{ float: "right", color: "#1890ff" }}
            href="/register"
            onClick={handleRegister}
          >
            Don't have an account
          </a>
        </div>
      </Form.Item>
      <Form.Item>
        <Button type="primary" htmlType="submit" style={{ width: "100%" }}>
          Login
        </Button>
      </Form.Item>
    </LoginForm>
  );
};
