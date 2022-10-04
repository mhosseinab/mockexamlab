import React, { FC } from "react";

import style from "styles/exams/questions.module.scss";
import ExamsFooter from "components/exams/Footer";
import { RootStateOrAny, useSelector } from "react-redux";
import TestHeader from "components/exams/TestHeader";
import { useRouter } from "next/router";

interface Props {
  children: React.ReactNode;
  title: string;
  playAudioHandler?: () => void;
}

const QuestionContainer: FC<Props> = ({ children }) => {
  const { tests, auth } = useSelector((state: RootStateOrAny) => state);
  const { query } = useRouter();
  return (
    <div className={style.question__layout}>
      {/*if only in listening*/}
      <TestHeader title={auth.userName} />
      <div className={style.children}>{children}</div>
      {query.level != "intro" && <ExamsFooter sections={tests.test.sections} />}
    </div>
  );
};
export default QuestionContainer;
