import DashboardLayout from "components/dashboard/DashboardLayout";
import ExamResults from "components/dashboard/exam/ExamResults";
import Header from "components/dashboard/exam/Header";
import Search from "components/dashboard/exam/Search";
import SearchResults from "components/dashboard/exam/SearchResults";
import React from "react";

const ExamContainer = () => {
  const [searchValue, setSearchValue] = React.useState("");
  const [currentCategory, setCurrentCategory] = React.useState("");
  return (
    <DashboardLayout>
      <Header />
      <Search value={searchValue} setValue={setSearchValue} />
      <SearchResults
        category={currentCategory}
        setCategory={setCurrentCategory}
      />
      <ExamResults value={searchValue} />

    </DashboardLayout>
  );
};

export default ExamContainer;
