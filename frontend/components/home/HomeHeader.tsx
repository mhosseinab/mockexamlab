import Image from "next/image";
import Link from "next/link";
import React from "react";
import style from "styles/home.module.scss";

const HomeHeader = () => {
  return (
    <div className={style.header__container}>
      <div className={style.info}>
        <h1>
          Find you path <br /> through learning
        </h1>
        <h4>
          We help you to find the best courses <br /> and virtual exam to learn
          anything!
        </h4>
        <div className={style.button_actions}>
          <Link href="">
            <a className={style.start}>
              Get started
              <svg
                width="24"
                height="24"
                viewBox="0 0 24 24"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M5 12H19"
                  stroke="#ffffff"
                  strokeWidth={2}
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
                <path
                  d="M12 5L19 12L12 19"
                  stroke="#ffffff"
                  strokeWidth={2}
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
              </svg>
            </a>
          </Link>
          <Link href="">
            <a className={style.more}>Learn more</a>
          </Link>
        </div>
        <div className={style.graph}>
          {activity.map((item, index) => (
            <div key={index} className={style.graph__item}>
              <span>{item.count} +</span>
              <h6>{item.title}</h6>
            </div>
          ))}
        </div>
        <div className={style.details}>
          <span>Hot</span>
          <p>Over 1000 students are taking the IELTS course and Exams</p>
        </div>
      </div>
      <div className={style.icons}>
        <div className={style.calender}>
          <Image
            src={"/img/calandar.png"}
            alt="calender"
            width={140}
            height={140}
          />
        </div>
        <div className={style.cap}>
          <Image src={"/img/cap.png"} alt="cap" width={295} height={250} />
        </div>
        <div className={style.file}>
          <Image src={"/img/file.png"} alt="file" width={140} height={140} />
        </div>
      </div>
    </div>
  );
};

const activity = [
  {
    title: "Courses",
    count: "1k",
  },
  {
    title: "Students",
    count: "1500",
  },
  {
    title: "Exam",
    count: "120",
  },
];

export default HomeHeader;
