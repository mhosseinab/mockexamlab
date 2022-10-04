import DashboardHeader from "components/ui/DashboardHeader";
import React, { FC } from "react";
import DashboardLayout from "../DashboardLayout";

interface Props {
  title: string;
}

const ProductDetails: FC<Props> = ({ title }) => {
  return (
    <DashboardLayout>
      <DashboardHeader title={title} />
    </DashboardLayout>
  );
};

export default ProductDetails;
