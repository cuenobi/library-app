"use client";
import React from "react";
import { Layout } from "antd";
import InvalidCredential from "@/components/Result/InvalidCredential";

const { Content } = Layout;

export default function Dashboard() {
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
          <InvalidCredential />
        </div>
      </Content>
    </Layout>
  );
}
