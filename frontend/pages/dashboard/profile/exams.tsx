import DashboardLayout from "components/dashboard/DashboardLayout";
import ExamResults from "components/dashboard/exam/ExamResults";
import Search from "components/dashboard/exam/Search";
import Header from "components/dashboard/profile/Header";
import DashboardHeader from "components/ui/DashboardHeader";
import Link from "next/link";
import React from "react";

const tags = [{ title: "Beta", color: "#FF7F37" }];

const Exams = () => {
  const [search, setSearch] = React.useState("");
  return (
    <DashboardLayout>
      <DashboardHeader title="Profile" />
      <Header tags={tags} />
      <Search setValue={setSearch} value={search} />
      <Link href="/dashboard/profile/createExam">
        <a className="full-width-blue">Create Exam</a>
      </Link>
      <ExamResults value={search} />
    </DashboardLayout>
  );
};

export default Exams;
