"use client";
import React from "react";
import { Menu, Avatar } from "antd";
import { UserOutlined, CodeOutlined, LogoutOutlined } from "@ant-design/icons";

interface RightMenuProps {
  mode: "vertical" | "horizontal";
}

const RightMenu: React.FC<RightMenuProps> = ({ mode }) => {
  const menuItems = [
    {
      key: "user-menu",
      label: (
        <div className="flex items-center space-x-2">
          <Avatar icon={<UserOutlined />} />
          <span className="username">Cue Sitthikorn</span>
        </div>
      ),
      children: [
        {
          key: "profile",
          label: (
            <>
              <UserOutlined /> Profile
            </>
          ),
        },
        {
          key: "logout",
          label: (
            <>
              <LogoutOutlined /> Logout
            </>
          ),
        },
      ],
    },
  ];

  return <Menu mode={mode} style={{ minWidth: "200px" }} items={menuItems} />;
};

export default RightMenu;