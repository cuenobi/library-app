"use client";

import React, { useState, useEffect } from "react";
import { Layout, Button, Drawer } from "antd";
import LeftMenu from "./LeftMenu";
import RightMenu from "./RightMenu";
import { MenuOutlined } from "@ant-design/icons";
import { usePathname } from "next/navigation";

const { Header } = Layout;

const Navbar: React.FC = () => {
  const [open, setOpen] = useState<boolean>(false);
  const pathname = usePathname();

  const toggleDrawer = () => {
    setOpen(!open);
  };

  useEffect(() => {
    setOpen(false);
  }, [pathname]);

  return (
    <Header className="fixed top-0 left-0 w-full z-50 bg-white shadow-md">
      <div className="flex justify-between items-center px-6">
        <a href="/" style={{ textDecoration: "none", color: "inherit" }}>
          <div className="flex items-center">
            <img src="./open-book.png" alt="" className="h-8 mr-2" />{" "}
            <h3 className="text-xl font-bold">Library</h3>{" "}
          </div>
        </a>

        <div className="hidden md:flex space-x-4">
          <LeftMenu mode="horizontal" />
          <RightMenu mode="horizontal" />
        </div>

        <Button className="md:hidden" type="text" onClick={toggleDrawer}>
          <MenuOutlined />
        </Button>
      </div>

      <Drawer
        title="Library"
        placement="right"
        closable
        onClose={toggleDrawer}
        open={open}
      >
        <LeftMenu mode="inline" />
        <RightMenu mode="vertical" />
      </Drawer>
    </Header>
  );
};

export default Navbar;
