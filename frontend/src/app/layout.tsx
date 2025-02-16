"use client";

import React from "react";
import { Layout } from "antd";
import Navbar from "../components/Navbar/NavBar";
import "../app/globals.css";

const { Content, Footer } = Layout;

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>
        <Layout
          style={{
            minHeight: "100vh",
            display: "flex",
            flexDirection: "column",
          }}
        >
          <Navbar />
          <Content style={{ padding: "0 48px", marginTop: 52, flex: 1 }}>
            {children}
          </Content>
          <Footer style={{ textAlign: "center" }}>
            Copyright ©{new Date().getFullYear()} Sitthikorn Khumthong
          </Footer>
        </Layout>
      </body>
    </html>
  );
}
