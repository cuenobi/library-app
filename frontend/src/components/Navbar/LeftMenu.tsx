// "use client";
// import React from "react";
// import { Menu } from "antd";
// import Link from "next/link";
// import { usePathname } from "next/navigation";

// interface LeftMenuProps {
//   mode: "vertical" | "horizontal" | "inline";
// }

// const LeftMenu: React.FC<LeftMenuProps> = ({ mode }) => {
//   const pathname = usePathname();

//   const items = [
//     { key: "/", label: <Link href="/">Home</Link> },
//     { key: "/register", label: <Link href="/register">Register</Link> },
//   ];

//   // return <Menu mode={mode} selectedKeys={[pathname || ""]} items={items} />;
//   return <Menu mode={mode} selectedKeys={[pathname || ""]} />;
// };

// export default LeftMenu;
