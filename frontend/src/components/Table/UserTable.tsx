import React, { useEffect, useState } from "react";
import { Table, Button, Space } from "antd";
import { SwapRightOutlined, SwapLeftOutlined } from "@ant-design/icons";
import axios from "axios";
import BorrowButton from "./BorrowButton";
import ReturnButton from "./ReturnButton";

const { Column } = Table;

interface UserData {
  key: string;
  created_at: string;
  username: string;
  name: string;
  borrowDetails: BorrowDetail[];
}

interface BorrowDetail {
  BookID: string;
  bookName: string;
  borrowedAt: number;
  ReturnedAt: number | null;
}

const App: React.FC = () => {
  const [data, setData] = useState<UserData[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string>("");

  const getAuthToken = () => localStorage.getItem("authToken");

  const fetchData = async () => {
    try {
      const token = getAuthToken();
      if (!token) {
        setError("No authorization token found.");
        setLoading(false);
        return;
      }

      const response = await axios.get("http://localhost:8080/users", {
        headers: { Authorization: `Bearer ${token}` },
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

  useEffect(() => {
    fetchData();
  }, []);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  const timeConverter = (UNIX_timestamp: number | null) => {
    if (!UNIX_timestamp || isNaN(UNIX_timestamp)) return "Not Available";

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
    return `${a.getDate().toString().padStart(2, "0")} ${
      months[a.getMonth()]
    } ${a.getFullYear()} ${a.getHours().toString().padStart(2, "0")}:${a
      .getMinutes()
      .toString()
      .padStart(2, "0")}:${a.getSeconds().toString().padStart(2, "0")}`;
  };

  return (
    <Table
      dataSource={data}
      expandable={{
        expandedRowRender: (record: UserData) => (
          <Table
            dataSource={record.borrowDetails}
            pagination={false}
            rowKey="bookName"
            style={{ background: "#f9f9f9", marginTop: 10 }}
          >
            <Column title="Book Name" dataIndex="BookName" key="bookName" />
            <Column
              title="Borrowed At"
              dataIndex="BorrowedAt"
              key="BorrowedAt"
              render={(borrowedAt: number | null) => timeConverter(borrowedAt)}
            />
            <Column
              title="Returned At"
              dataIndex="ReturnedAt"
              key="ReturnedAt"
              render={(returnedAt: number | null, borrowRecord: BorrowDetail) =>
                returnedAt === null ? (
                  <ReturnButton
                    userID={record.key}
                    key={borrowRecord.BookID}
                    bookID={borrowRecord.BookID}
                    onReturnSuccess={fetchData}
                  />
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
        title=""
        key="action"
        render={(_: any, record: UserData) => (
          <Space size="middle">
            <BorrowButton
              key={record.key}
              userID={record.key}
              onBorrowSuccess={fetchData}
            />
          </Space>
        )}
      />
    </Table>
  );
};

export default App;
