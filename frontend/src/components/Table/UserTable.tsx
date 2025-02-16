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
  borrowDetails: BorrowDetail[];
}

interface BorrowDetail {
  bookName: string;
  borrowedAt: number;
  ReturnedAt: number | null;
}

const App: React.FC = () => {
  const [data, setData] = useState<UserData[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>("");

  const getAuthToken = () => {
    return localStorage.getItem("authToken");
  };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const token = getAuthToken();
        if (!token) {
          setError("No authorization token found.");
          setLoading(false);
          return;
        }

        const response = await axios.get("http://localhost:8080/users", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        console.log(response.data);
        const users = response.data.users.map((user: any) => ({
          key: user.ID,
          created_at: user.CreatedAt,
          username: user.Username,
          name: user.Name,
          borrowDetails: user.BorrowDetails || [],
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

  const timeConverter = (UNIX_timestamp: number | null) => {
    if (
      UNIX_timestamp === null ||
      UNIX_timestamp === undefined ||
      isNaN(UNIX_timestamp)
    ) {
      return "Not Available";
    }

    const a = new Date(UNIX_timestamp * 1000);
    const months = [
      "Jan",
      "Feb",
      "Mar",
      "Apr",
      "May",
      "Jun",
      "Jul",
      "Aug",
      "Sep",
      "Oct",
      "Nov",
      "Dec",
    ];
    const year = a.getFullYear();
    const month = months[a.getMonth()];
    const date = a.getDate();
    const hour = a.getHours();
    const min = a.getMinutes();
    const sec = a.getSeconds();

    const time = `${date} ${month} ${year} ${hour}:${min}:${sec}`;
    return time;
  };

  return (
    <Table
      dataSource={data}
      expandable={{
        expandedRowRender: (record) => (
          <Table
            dataSource={record.borrowDetails}
            pagination={false}
            rowKey="bookName"
            style={{ background: "#f9f9f9", marginTop: 10 }}
          >
            <Column title="Book Name" dataIndex="BookName" key="BookName" />
            <Column
              title="Borrowed At"
              dataIndex="BorrowedAt"
              key="BorrowedAt"
              render={(BorrowedAt: number | null) => timeConverter(BorrowedAt)}
            />
            <Column
              title="Returned At"
              dataIndex="returnedAt"
              key="returnedAt"
              render={(returnedAt: number | null) =>
                returnedAt === null || returnedAt === undefined ? (
                  <Button
                    type="primary"
                    shape="round"
                    icon={<SwapLeftOutlined />}
                  >
                    Return
                  </Button>
                ) : (
                  timeConverter(returnedAt)
                )
              }
            />
          </Table>
        ),
        rowExpandable: (data) => data.name !== "Not Expandable",
      }}
      rowKey="key"
    >
      <Column title="Created At" dataIndex="created_at" key="created_at" />
      <Column title="Username" dataIndex="username" key="username" />
      <Column title="Name" dataIndex="name" key="name" />
      <Column
        title="Action"
        key="action"
        render={(_: any, record: UserData) => (
          <Space size="middle">
            <Button type="primary" shape="round" icon={<SwapRightOutlined />}>
              Borrow
            </Button>
          </Space>
        )}
      />
    </Table>
  );
};

export default App;
