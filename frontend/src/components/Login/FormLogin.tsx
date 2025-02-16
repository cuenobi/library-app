// src/components/forms/LoginForm.tsx
import { LockOutlined, UserOutlined } from "@ant-design/icons";
import {
  LoginForm,
  ProConfigProvider,
  ProFormText,
} from "@ant-design/pro-components";
import { Button, Form, Input, Checkbox, message } from "antd";
import { useAuth } from "../../hooks/useAuth";

export default () => {
  const { login } = useAuth();

  const handleSubmit = async (values: {
    username: string;
    password: string;
  }) => {
    const { username, password } = values;

    const success = await login(username, password);
    if (success) {
      message.success("Login successful!");
      // ถ้าล็อกอินสำเร็จ, สามารถนำไปยังหน้าอื่นๆ เช่น dashboard
      // router.push("/dashboard");
    } else {
      message.error("Login failed! Please check your credentials.");
    }
  };

  return (
    <LoginForm
      logo="./group.png"
      title="Member"
      subTitle="Member Register"
      name="login"
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
        <Checkbox>Remember me</Checkbox>
      </Form.Item>
      <Form.Item>
        <Button type="primary" htmlType="submit" style={{ width: "100%" }}>
          Login
        </Button>
      </Form.Item>
    </LoginForm>
  );
};
