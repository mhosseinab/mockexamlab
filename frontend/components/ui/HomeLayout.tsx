import React, { FC } from "react";
import style from "styles/layout.module.scss";

interface Props {
  children: React.ReactNode;
}

const HomeLayout: FC<Props> = ({ children }) => {
  return <div className={style.container}>{children}</div>;
};

export default HomeLayout;
