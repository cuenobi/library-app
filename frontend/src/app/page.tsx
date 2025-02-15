"use client";
import React from "react";
import { Layout } from "antd";
import FormLogin from "@/components/Login/FormLogin";

const { Content } = Layout;

export default function Page() {
  return (
    <Layout>
      <Content style={{ padding: "24px 48px" }}>
        <div
          style={{
            background: "white",
            padding: 24,
            minHeight: 280,
            borderRadius: 8,
          }}
        >
          <FormLogin/>
        </div>
      </Content>
    </Layout>
  );
}
