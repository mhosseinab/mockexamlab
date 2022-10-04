import React, { FC } from "react";

import style from "styles/exams/questions.module.scss";
import { useRouter } from "next/router";
import TestHeader from "./TestHeader";

interface Props {
  title: string;
  time: string;
  name: string;
  section: string;
}

const Instruction: FC<Props> = ({ title, time, name, section }) => {
  const { pathname, push } = useRouter();
  return (
    <div className={style.exam__intro}>
      <div className={style.title}>
        <h4>{title}</h4>
        <div className={style.br} />
      </div>
      <div className={style.menu}>
        <div className={style.head}>
          <h4>{name}</h4>
          <h4>Time : Approximately {time} Minutes</h4>
        </div>
        <ul>
          <h4>INSTRUCTIONS TO CANDIDATES</h4>
          <li>You can change your answers at any time during the test.</li>
          <li>INFORMATION FOR CANDIDATES</li>
          <li>There are 40 questions in this test.</li>
          <li>Each question carries one mark.</li>
          <li>There are four parts to the test</li>
        </ul>
      </div>
      <button
        className={style.start}
        onClick={() => {
          push({
            pathname,
            query: {
              section: section,
              level: "questions",
            },
          });
        }}
      >
        Start test
      </button>
    </div>
  );
};
export default Instruction;
