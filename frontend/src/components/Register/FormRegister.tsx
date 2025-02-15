import { LockOutlined, UserOutlined, IdcardOutlined } from "@ant-design/icons";
import {
  LoginForm,
  ProConfigProvider,
  ProFormText,
} from "@ant-design/pro-components";
import { theme, Button, ConfigProvider } from "antd";
import { useForm } from "antd/es/form/Form";

export default () => {
  const { token } = theme.useToken();
  const [form] = useForm();

  const handleSubmit = (values: any) => {
    console.log(values);
  };

  return (
    <ConfigProvider theme={{ token }}>
      <ProConfigProvider hashed={false}>
        <div style={{ backgroundColor: token.colorBgContainer }}>
          <LoginForm
            logo="./group.png"
            title="Member"
            subTitle="Member Register"
            onFinish={handleSubmit}
            submitter={false}
            form={form}
          >
            <ProFormText
              name="username"
              fieldProps={{
                size: "large",
                prefix: <UserOutlined className={"prefixIcon"} />,
                autoFocus: true,
              }}
              placeholder={"username"}
              rules={[
                {
                  required: true,
                  message: "Please enter your username!",
                },
              ]}
            />
            <ProFormText.Password
              name="password"
              fieldProps={{
                size: "large",
                prefix: <LockOutlined className={"prefixIcon"} />,
              }}
              placeholder={"password"}
              rules={[
                {
                  required: true,
                  message: "Please enter your password!",
                },
              ]}
            />
            <ProFormText.Password
              name="re-password"
              fieldProps={{
                size: "large",
                prefix: <LockOutlined className={"prefixIcon"} />,
              }}
              placeholder={"re-password"}
              rules={[
                {
                  required: true,
                  message: "Please enter your re-password!",
                },
                {
                  validator: async (_, value) => {
                    const password = form.getFieldValue("password");
                    if (value && value !== password) {
                      return Promise.reject(
                        new Error("Passwords do not match!")
                      );
                    }
                    return Promise.resolve();
                  },
                },
              ]}
            />
            <ProFormText
              name="Name"
              fieldProps={{
                size: "large",
                prefix: <IdcardOutlined className={"prefixIcon"} />,
                autoFocus: true,
              }}
              placeholder={"full name"}
              rules={[
                {
                  required: true,
                  message: "Please enter your Name!",
                },
              ]}
            />
            <Button
              type="primary"
              htmlType="submit"
              style={{ width: "100%", marginTop: 24 }}
            >
              Submit
            </Button>
          </LoginForm>
        </div>
      </ProConfigProvider>
    </ConfigProvider>
  );
};
