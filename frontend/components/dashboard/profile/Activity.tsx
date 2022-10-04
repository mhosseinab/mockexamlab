import Image from "next/image";
import React from "react";

import style from "styles/dashboard/profile.module.scss";

const Activity = () => {
  return (
    <div className={style.profile__activity}>
      <h1>Latest Activity</h1>
      {activities.map((activity, index) => (
        <div key={index} className={style.profile__activity_item}>
          <span className={style.icon}>
            <Image
              src={activity.icon}
              alt={activity.title}
              width={14}
              height={14}
            />
          </span>
          <h4>{activity.title}</h4>
        </div>
      ))}
    </div>
  );
};

const activities = [
  { title: "Loged in in 21/8", icon: "/img/icons/person.svg" },
  { title: "Take an IELTS Exam", icon: "/img/icons/bookmark.svg" },
  { title: "Add MohsenKhani22 As a friend", icon: "/img/icons/heart.svg" },
  { title: "Loged in in 21/8", icon: "/img/icons/person.svg" },
];

export default Activity;
