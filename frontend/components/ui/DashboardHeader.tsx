import Image from "next/image";
import Link from "next/link";
import React, { FC } from "react";

import style from "styles/dashboard/dashboard.module.scss";

interface Props {
  title: string;
}

const DashboardHeader: FC<Props> = ({ title }) => {
  return (
    <div className={style.dashboard__header}>
      <h3>{title}</h3>
      <div className={style.action}>
        <Link href="/dashboard">
          <span className={style.icon}>
            <Image
              src="/img/icons/person.svg"
              alt="person icon"
              width={14}
              height={20}
            />
          </span>
        </Link>
        <span className={style.icon}>
          <Image
            src="/img/icons/notifaction.svg"
            alt="notifaction icon"
            width={18}
            height={21}
          />
        </span>
      </div>
    </div>
  );
};

export default DashboardHeader;
