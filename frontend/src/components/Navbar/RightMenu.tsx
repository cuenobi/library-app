"use client";
import React, { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { Menu, Avatar } from "antd";
import { UserOutlined, LogoutOutlined } from "@ant-design/icons";

interface RightMenuProps {
  mode: "vertical" | "horizontal";
}

const RightMenu: React.FC<RightMenuProps> = ({ mode }) => {
  const router = useRouter();
  const [username, setUsername] = useState<string | null>(null);
  
  useEffect(() => {
    setUsername(localStorage.getItem("username") || null);

    const handleStorageChange = () => {
      setUsername(localStorage.getItem("username") || null);
    };

    window.addEventListener("storage", handleStorageChange);
    return () => {
      window.removeEventListener("storage", handleStorageChange);
    };
  }, []);

  const handleLogout = () => {
    localStorage.removeItem("authToken");
    localStorage.removeItem("username");
    window.dispatchEvent(new Event("storage"));

    router.push("/");
  };

  const menuItems = [
    {
      key: "user-menu",
      label: (
        <div className="flex items-center space-x-2">
          <Avatar icon={<UserOutlined />} />
          <span className="username">{username || "Guest"}</span>{" "}
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
            <div onClick={handleLogout}>
              <LogoutOutlined /> Logout
            </div>
          ),
        },
      ],
    },
  ];

  return <Menu mode={mode} style={{ minWidth: "200px" }} items={menuItems} />;
};

export default RightMenu;
