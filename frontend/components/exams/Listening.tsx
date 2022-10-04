import React, { FC } from "react";
import CheckQuestion from "./CheckQuestion";
import { RootStateOrAny, useSelector } from "react-redux";
import { useRouter } from "next/router";
import Instruction from "./Instruction";
import style from "styles/exams/questions.module.scss";
import QuestionContainer from "../../container/QuessionContainer";
import NoteCompletion from "./NoteCompletion";

interface Props {}

const Listening: FC<Props> = ({}) => {
  const { tests, auth } = useSelector((state: RootStateOrAny) => state);
  const { query } = useRouter();

  if (query.level == "intro")
    return (
      <QuestionContainer title={""}>
        <Instruction
          title={"Instruction"}
          time={"30"}
          name={"IELTS Reading"}
          section={"listening"}
        />
      </QuestionContainer>
    );
  const questions =
    tests?.test?.sections[1]?.questionGroups[1]?.multipleChoice || null;
  const noteCompletionQuestions = tests?.test?.sections[1]?.questionGroups[0];
  return (
    <QuestionContainer title={auth.userName}>
      <div className={style.exam__title}>
        <h4>{tests.test.sections[2].componentType}</h4>
        <div className={style.br} />
      </div>
      <h4 className={style.exam__tests}>
        Questions {"  "}
        {questions[0].q_number}- {questions[0].q_number + questions.length}
      </h4>
      <p className={style.section__details}>
        {tests?.test?.sections[1]?.questionGroups[1]?.description}
      </p>
      {questions?.map((item: any, index: number) => (
        <CheckQuestion
          key={index}
          title={item.q_number + " - " + item.title}
          question={item.answers}
        />
      ))}
      <h1>{noteCompletionQuestions.title}</h1>

      <p style={{ margin: "1em 0" }}>{noteCompletionQuestions.description}</p>
      <NoteCompletion question={noteCompletionQuestions.noteCompletion.note} />
    </QuestionContainer>
  );
};

export default Listening;
