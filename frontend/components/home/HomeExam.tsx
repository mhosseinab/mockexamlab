/* eslint-disable @next/next/no-img-element */
import React from "react";
import style from "styles/home.module.scss";

const HomeExam = () => {
  return (
    <div className={style.exam__container}>
      <div className={style.container}>
        <div className={style.info}>
          <h1>
            A HUGE <br />
            database of exams!
          </h1>
          <h3>We are hosting a huge database of exams that you can take!</h3>
        </div>
        <div className={style.actions}>
          {exams.map((item, index) => (
            <span key={index}>
              <div className={style.bookmark__container}>
                <svg
                  className={style.bookmark}
                  width="12"
                  height="16"
                  viewBox="0 0 12 16"
                  fill="none"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    d="M10.5 0.5H1.5C0.675 0.5 0 1.175 0 2V14C0 14.825 0.675 15.5 1.5 15.5H10.5C11.325 15.5 12 14.825 12 14V2C12 1.175 11.325 0.5 10.5 0.5ZM1.5 2H5.25V8L3.375 6.875L1.5 8V2Z"
                    fill="white"
                  />
                </svg>
              </div>

              {item}
              <svg
                width="24"
                height="24"
                viewBox="0 0 24 24"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  className={style.arrow}
                  d="M5 12H19"
                  stroke="#52C3FF"
                  strokeWidth={2}
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
                <path
                  d="M12 5L19 12L12 19"
                  stroke="#52C3FF"
                  className={style.arrow}
                  strokeWidth={2}
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
              </svg>
            </span>
          ))}
        </div>
      </div>
    </div>
  );
};

const exams = ["IELTS", "HSD B1", "TOEFL", "GEBT"];

export default HomeExam;
