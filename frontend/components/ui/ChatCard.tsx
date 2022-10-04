import Image from "next/image";
import React, { FC } from "react";
import style from "styles/dashboard/profile.module.scss";

interface Props {
  image: string;
  description: string;
}

const ChatCard: FC<Props> = ({ description, image }) => {
  return (
    <div className={style.chat__card}>
      <Image src={image} alt="profile" width={32} height={32} />
      <p className={style.chat}>{description}</p>
    </div>
  );
};

export default ChatCard;
