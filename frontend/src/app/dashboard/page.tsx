"use client";
import React, { useEffect } from "react";
import { useRouter } from "next/navigation";
import { Layout } from "antd";
import UserTable from "@/components/Table/UserTable";

const { Content } = Layout;

export default function Dashboard() {
  const router = useRouter();

  useEffect(() => {
    const role = localStorage.getItem("role");

    if (role !== "2") {
      router.push("/error");
    }
  }, []);

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
          <UserTable />
        </div>
      </Content>
    </Layout>
  );
}
