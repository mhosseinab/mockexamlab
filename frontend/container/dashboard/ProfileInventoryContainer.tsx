import DashboardLayout from "components/dashboard/DashboardLayout";
import CardList from "components/dashboard/profile/CardList";
import Header from "components/dashboard/profile/Header";
import InventoryResults from "components/dashboard/profile/InventoryResults";
import DashboardHeader from "components/ui/DashboardHeader";
import React from "react";

const ProfileInventoryContainer = () => {
  const [currentMediaType, setCurrentMediaType] = React.useState<string>("Event");

  return (
    <DashboardLayout>
      <DashboardHeader title="Profile" />
      <Header tags={tags} />
      <span className="full-width-light">
        <svg
          width="16"
          height="20"
          viewBox="0 0 16 20"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            fillRule="evenodd"
            clipRule="evenodd"
            d="M6 0C4.89543 0 4 0.89543 4 2V3H2C0.89543 3 0 3.89543 0 5V17C0 18.1046 0.895431 19 2 19H3C3 19.5523 3.44772 20 4 20C4.55228 20 5 19.5523 5 19H11C11 19.5523 11.4477 20 12 20C12.5523 20 13 19.5523 13 19H14C15.1046 19 16 18.1046 16 17V5C16 3.89543 15.1046 3 14 3H12V2C12 0.89543 11.1046 0 10 0H6ZM12 17H14V5H10H6L2 5V17H4H12ZM10 3V2H6V3H10ZM6 8C6.55228 8 7 8.44771 7 9V13C7 13.5523 6.55228 14 6 14C5.44772 14 5 13.5523 5 13V9C5 8.44771 5.44772 8 6 8ZM11 9C11 8.44771 10.5523 8 10 8C9.44771 8 9 8.44771 9 9V13C9 13.5523 9.44771 14 10 14C10.5523 14 11 13.5523 11 13V9Z"
            fill="#52C3FF"
          />
        </svg>
        Inventory
      </span>
      <InventoryResults
        currentMediaType={currentMediaType}
        setCurrentMediaType={setCurrentMediaType}
      />
      <CardList title={currentMediaType} link="" items={eventList} />
    </DashboardLayout>
  );
};

const tags = [
  { title: "New", color: "#3ac922" },
  { title: "Super Active", color: "#3ac922" },
  { title: "-A- Grade", color: "#3ac922" },
];

const eventList = [
  { title: "Event Name", image: "/img/dashboard/profile-banner.png", link: "/" },
  { title: "Event Name", image: "/img/dashboard/profile-banner.png", link: "/" },
  { title: "Event Name", image: "/img/dashboard/profile-banner.png", link: "/" },
  { title: "Event Name", image: "/img/dashboard/profile-banner.png", link: "/" },
];

export default ProfileInventoryContainer;
