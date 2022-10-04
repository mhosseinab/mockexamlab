import DashboardLayout from "components/dashboard/DashboardLayout";
import Search from "components/dashboard/exam/Search";
import CardList from "components/dashboard/profile/CardList";
import Header from "components/dashboard/shop/Header";
import Routing from "components/dashboard/shop/Routing";
import DashboardHeader from "components/ui/DashboardHeader";
import React from "react";

const BooksContainer = () => {
  const [search, setSearch] = React.useState("");
  const [category, setCategory] = React.useState("");
  return (
    <DashboardLayout>
      <DashboardHeader title="Books" />
      <Header image="/img/dashboard/books.png" link="/" title="Check Latest" width="164" height="164" />
      <Search setValue={setSearch} value={search} />
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
export default BooksContainer;
