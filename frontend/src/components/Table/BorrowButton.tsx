import React, { useState } from "react";
import { Button, Flex, Modal } from "antd";
import { SwapRightOutlined, SwapLeftOutlined } from "@ant-design/icons";
import BooksTable from "./BookTable";

interface Props {
  userID: string;
  onBorrowSuccess: () => void;
}

const App: React.FC<Props> = ({ userID, onBorrowSuccess }) => {
  const [open, setOpen] = useState(false);
  const [openResponsive, setOpenResponsive] = useState(false);

  const handleBorrowSuccess = () => {
    setOpenResponsive(false);
    onBorrowSuccess();
  };

  return (
    <Flex vertical gap="middle" align="flex-start">
      {/* Responsive */}
      <Button
        type="primary"
        shape="round"
        icon={<SwapRightOutlined />}
        onClick={() => setOpenResponsive(true)}
      >
        Borrow
      </Button>
      <Modal
        title="All books"
        centered
        open={openResponsive}
        onOk={() => handleBorrowSuccess()}
        onCancel={() => setOpenResponsive(false)}
        width={{
          xs: "90%",
          sm: "80%",
          md: "70%",
          lg: "60%",
          xl: "50%",
          xxl: "40%",
        }}
      >
        <BooksTable userID={userID} />
      </Modal>
    </Flex>
  );
};

export default App;
