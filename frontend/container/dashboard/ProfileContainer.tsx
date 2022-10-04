import DashboardLayout from "components/dashboard/DashboardLayout";
import Activity from "components/dashboard/profile/Activity";
import Header from "components/dashboard/profile/Header";
import Inputs from "components/dashboard/profile/Inputs";
import Routing from "components/dashboard/profile/Routing";
import DashboardHeader from "components/ui/DashboardHeader";
import React from "react";

const tags = [
  { title: "New", color: "#3ac922" },
  { title: "Super Active", color: "#3ac922" },
  { title: "-A- Grade", color: "#3ac922" },
];

const ProfileContainer = () => {
  return (
    <DashboardLayout>
      <DashboardHeader title="Profile" />
      <Header tags={tags} />
      <Inputs />
      <Routing />
      <Activity />
    </DashboardLayout>
  );
};

export default ProfileContainer;
