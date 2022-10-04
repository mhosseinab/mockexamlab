import DashboardLayout from "components/dashboard/DashboardLayout";
import AddCard from "components/dashboard/financial/AddCard";
import Card from "components/dashboard/financial/Card";
import Payments from "components/dashboard/financial/Payments";
import DashboardHeader from "components/ui/DashboardHeader";
import React from "react";

import style from "styles/dashboard/financial.module.scss";

const FinancialContainer = () => {
  const [card, setCard] = React.useState<string>("");
  return (
    <DashboardLayout>
      <DashboardHeader title="Wallet" />
      <div className={style.card__container}>
        <Card />
        <AddCard setValue={setCard} value={card} />
      </div>
      <Payments />
    </DashboardLayout>
  );
};

export default FinancialContainer;
