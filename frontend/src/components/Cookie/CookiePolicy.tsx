"use client";
import { useEffect, useState } from "react";
import { Button } from "antd";
import { motion } from "framer-motion";

const COOKIE_ACCEPTED_KEY = "cookieAccepted";

const CookiePolicy: React.FC = () => {
  const [visible, setVisible] = useState<boolean>(false);

  useEffect(() => {
    const hasAccepted = localStorage.getItem(COOKIE_ACCEPTED_KEY);
    if (!hasAccepted) {
      setVisible(true);
    }
  }, []);

  const handleAccept = () => {
    localStorage.setItem(COOKIE_ACCEPTED_KEY, "true");
    setVisible(false);
  };

  if (!visible) return null;

  return (
    <motion.div
      initial={{ opacity: 0, y: 30 }}
      animate={{ opacity: 1, y: 0 }}
      exit={{ opacity: 0, y: 30 }}
      transition={{ duration: 0.4, ease: "easeInOut" }}
      className="fixed bottom-6 left-6 right-6 md:left-auto md:right-6 max-w-md bg-white/80 dark:bg-gray-900/90 backdrop-blur-md p-5 shadow-xl rounded-xl border border-gray-200 dark:border-gray-800 flex flex-col space-y-3 z-50"
    >
      <h3 className="text-lg font-semibold text-gray-900 dark:text-white">
        🍪 Cookie Usage
      </h3>
      <p className="text-sm text-gray-600 dark:text-gray-300">
        Our website uses cookies to enhance your browsing experience.
        Read more at{" "}
        <a
          href="/privacy-policy"
          className="text-blue-600 dark:text-blue-400 underline"
        >
          Privacy Policy
        </a>
      </p>
      <div className="flex justify-end space-x-3">
        <Button
          size="middle"
          onClick={handleAccept}
          className="text-gray-700 dark:text-gray-300 border-gray-300 dark:border-gray-700"
        >
          Decline
        </Button>
        <Button
          type="primary"
          size="middle"
          onClick={handleAccept}
          className="bg-blue-600 hover:bg-blue-700"
        >
          Accept
        </Button>
      </div>
    </motion.div>
  );
};

export default CookiePolicy;
