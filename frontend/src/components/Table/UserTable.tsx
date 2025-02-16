import React, { useEffect, useState } from "react";
import { Table, Button, Space } from "antd";
import { SwapRightOutlined, SwapLeftOutlined } from "@ant-design/icons";
import axios from "axios";

const { Column } = Table;

interface UserData {
  key: string;
  created_at: string;
  username: string;
  name: string;
}

const App: React.FC = () => {
  const [data, setData] = useState<UserData[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>("");

  // ฟังก์ชันนี้จะดึง token จาก localStorage
  const getAuthToken = () => {
    return localStorage.getItem("authToken");
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const token = getAuthToken(); // ดึง token
        if (!token) {
          setError("No authorization token found.");
          setLoading(false);
          return;
        }

        // ส่ง request พร้อมกับ token ใน header
        const response = await axios.get("http://localhost:8080/users", {
          headers: {
            Authorization: `Bearer ${token}`, // เพิ่ม Bearer token ใน header
          },
        });

        const users = response.data.users.map((user: any) => ({
          key: user.ID, // ใช้ `ID` เป็น key
          created_at: user.CreatedAt, // ใช้ `CreatedAt` ในการแสดงวันที่
          username: user.Username, // ใช้ `Username`
          name: user.Name, // ใช้ `Name`
        }));

        setData(users);
        setLoading(false);
      } catch (err) {
        setError("Failed to fetch users");
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  return (
    <Table dataSource={data}>
      <Column title="Created At" dataIndex="created_at" key="created_at" />
      <Column title="Username" dataIndex="username" key="username" />
      <Column title="Name" dataIndex="name" key="name" />
      <Column
        title="Action"
        key="action"
        render={(_: any, record: UserData) => (
          <Space size="middle">
            <Button
              type="primary"
              shape="round"
              icon={<SwapRightOutlined />}
            >
              Borrow
            </Button>
            <Button
              type="primary"
              shape="round"
              icon={<SwapLeftOutlined />}
            >
              Return
            </Button>
          </Space>
        )}
      />
    </Table>
  );
};

export default App;
