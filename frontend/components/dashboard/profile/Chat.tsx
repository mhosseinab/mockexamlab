import ChatCard from "components/ui/ChatCard";
import DashboardHeader from "components/ui/DashboardHeader";
import Image from "next/image";
import React from "react";

import style from "styles/dashboard/profile.module.scss";
import DashboardLayout from "../DashboardLayout";

const Chat = () => {
  return (
    <DashboardLayout>
      <div className={style.chat__container}>
        <DashboardHeader title="Chat with Cna" />
        <ChatCard
          description="Amet minim mollit non deserunt ullamco est sit aliqua dolor do amet sint. Velit officia consequat duis enim velit mollit. Exercitation veniam consequat sunt."
          image="/img/dashboard/profile-pic.png"
        />
      </div>
    </DashboardLayout>
  );
};

export default Chat;
