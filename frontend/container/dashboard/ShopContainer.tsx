import DashboardLayout from "components/dashboard/DashboardLayout";
import Search from "components/dashboard/exam/Search";
import CardList from "components/dashboard/profile/CardList";
import Header from "components/dashboard/shop/Header";
import Routing from "components/dashboard/shop/Routing";
import DashboardHeader from "components/ui/DashboardHeader";
import React from "react";

const ShopContainer = () => {
  const [search, setSearch] = React.useState("");
  const [category, setCategory] = React.useState("");
  return (
    <DashboardLayout>
      <DashboardHeader title="Shop" />
      <Header image="/img/dashboard/trophy.png" link="/dashboard/shop/books" title="Lets get" br="premium" height="208" width="208" />
      <Search setValue={setSearch} value={search} />
      <Routing category={category} setCategory={setCategory} />
      <CardList items={categoryList} title="Products" link="" />
    </DashboardLayout>
  );
};

const categoryList = [
  { title: "Book Name", image: "/img/dashboard/profile-banner.png", price: 25 },
  { title: "Book Name", image: "/img/dashboard/profile-banner.png", price: 25 },
  { title: "Book Name", image: "/img/dashboard/profile-banner.png", price: 25 },
  { title: "Book Name", image: "/img/dashboard/profile-banner.png", price: 25 },
];
export default ShopContainer;
