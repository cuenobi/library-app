import React, { useEffect, useState } from "react";
import { Button, Space, Table, Tag, message, Modal } from "antd";
import { SelectOutlined, ExclamationCircleOutlined } from "@ant-design/icons";
import type { TableProps } from "antd";

interface Book {
  key: string;
  name: string;
  category: string;
  status: string;
  stock: number;
}

interface BooksTableProps {
  userID: string; // รับ userID จาก parent
}

const BooksTable: React.FC<BooksTableProps> = ({ userID }) => {
  const [books, setBooks] = useState<Book[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [selectedBook, setSelectedBook] = useState<Book | null>(null);
  const [modalVisible, setModalVisible] = useState<boolean>(false);

  useEffect(() => {
    fetchBooks();
  }, []);

  const fetchBooks = async () => {
    setLoading(true);
    try {
      const token = localStorage.getItem("authToken");
      if (!token) {
        throw new Error("No authentication token found");
      }

      const response = await fetch("http://localhost:8080/books", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        throw new Error("Failed to fetch books");
      }

      const result = await response.json();
      const formattedBooks = result.books.map((book: any) => ({
        key: book.ID,
        name: book.Name,
        category: book.Category,
        status: book.Status,
        stock: book.Stock,
      }));

      setBooks(formattedBooks);
    } catch (error) {
      message.error("Error fetching books data");
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const handleCheckout = async () => {
    if (!selectedBook) return;

    try {
      const token = localStorage.getItem("authToken");
      if (!token) {
        throw new Error("No authentication token found");
      }

      const response = await fetch(
        `http://localhost:8080/book/borrow/${userID}/${selectedBook.key}`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
          body: JSON.stringify({
            
          })
        }
      );

      if (!response.ok) {
        throw new Error("Failed to checkout book");
      }

      message.success(`Successfully borrowed "${selectedBook.name}"`);
      setModalVisible(false);
      fetchBooks(); // Refresh data after checkout
    } catch (error) {
      message.error("Error processing checkout");
      console.error(error);
    }
  };

  const showConfirmModal = (book: Book) => {
    setSelectedBook(book);
    setModalVisible(true);
  };

  const columns: TableProps<Book>["columns"] = [
    {
      title: "Name",
      dataIndex: "name",
      key: "name",
      render: (text) => <a>{text}</a>,
    },
    {
      title: "Category",
      dataIndex: "category",
      key: "category",
    },
    {
      title: "Status",
      dataIndex: "status",
      key: "status",
      render: (status) => {
        const statusMapping: Record<string, string> = {
          "01": "Available",
          "02": "Checked Out",
        };
        return (
          <Tag color={status === "01" ? "green" : "red"}>
            {statusMapping[status] || "Unknown"}
          </Tag>
        );
      },
    },
    {
      title: "Stock",
      dataIndex: "stock",
      key: "stock",
    },
    {
      title: "",
      key: "action",
      render: (_, record) => (
        <Space size="middle">
          <Button
            type="default"
            shape="round"
            icon={<SelectOutlined />}
            disabled={record.status !== "01"}
            onClick={() => showConfirmModal(record)}
          >
            Checkout
          </Button>
        </Space>
      ),
    },
  ];

  return (
    <>
      <Table<Book> columns={columns} dataSource={books} loading={loading} />

      <Modal
        title="Confirm Checkout"
        visible={modalVisible}
        onOk={handleCheckout}
        onCancel={() => setModalVisible(false)}
        okText="Confirm"
        cancelText="Cancel"
      >
        {selectedBook && (
          <p>Are you sure you want to borrow "{selectedBook.name}"?</p>
        )}
      </Modal>
    </>
  );
};

export default BooksTable;
