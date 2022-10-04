import React, { FC } from "react";
import Sidebar from "./Sidebar";
import style from "styles/dashboard/dashboard.module.scss";
import MobileMenu from "components/ui/MobileMenu";

interface Props {
  children: React.ReactNode;
}

const DashboardLayout: FC<Props> = ({ children }) => {
  return (
    <div className={style.layout}>
      <Sidebar />
      <MobileMenu/>
      <div className={style.children}>{children}</div>
    </div>
  );
};

export default DashboardLayout;
