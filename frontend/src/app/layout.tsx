"use client";

import React from "react";
import { Layout } from "antd";
import Navbar from "../components/Navbar/NavBar";
import "../app/globals.css";
import CookiePolicy from "../components/Cookie/CookiePolicy";

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
            Copyright ©{new Date().getFullYear()}{" "}
            <a
              href="https://www.linkedin.com/in/sitthikorn-khu/"
              target="_blank"
            >
              Sitthikorn Khumthong
            </a>
          </Footer>
          <CookiePolicy />
        </Layout>
      </body>
    </html>
  );
}
