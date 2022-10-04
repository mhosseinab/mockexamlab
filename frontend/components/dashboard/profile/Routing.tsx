import Link from "next/link";
import React from "react";

import style from "styles/dashboard/profile.module.scss";

const Routing = () => {
  return (
    <div className={style.profile__routing}>
      <Link href="/dashboard/profile/exams">
        <a className={style.exam}>My Exams</a>
      </Link>
      <Link href="/">
        <a>My Routs</a>
      </Link>
      <div className={style.user__routes}>
        <Link href="/dashboard/stats">
          <a>
            Stats
            <svg
              width="20"
              height="20"
              viewBox="0 0 20 20"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M3 4C2.44772 4 2 3.55228 2 3C2 2.44772 2.44772 2 3 2C3.55228 2 4 2.44772 4 3C4 3.55228 3.55228 4 3 4ZM2.02723e-07 3C1.82965e-07 4.65685 1.34315 6 3 6C4.30622 6 5.41746 5.16519 5.82929 4L10 4C11.9921 4 13 5.34562 13 6.5C13 7.19482 12.8022 7.7952 12.4009 8.21659C12.0112 8.62574 11.3022 9 10 9C8.30216 9 7.0112 9.50074 6.15086 10.4041C5.30219 11.2952 5 12.4448 5 13.5C5 15.8456 6.99206 18 10 18L17 18L17 20L20 17L17 14L17 16L10 16C8.00794 16 7 14.6544 7 13.5C7 12.8052 7.19781 12.2048 7.59914 11.7834C7.9888 11.3743 8.69784 11 10 11C11.6978 11 12.9888 10.4993 13.8491 9.59591C14.6978 8.7048 15 7.55518 15 6.5C15 4.15438 13.0079 2 10 2L5.82929 2C5.41746 0.834808 4.30622 5.13511e-08 3 3.57746e-08C1.34315 1.60169e-08 2.22481e-07 1.34315 2.02723e-07 3Z"
                fill="#52C3FF"
              />
            </svg>
          </a>
        </Link>
        <Link href="/dashboard/profile/inventory">
          <a>
            Inventory
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
          </a>
        </Link>
      </div>
    </div>
  );
};

export default Routing;
