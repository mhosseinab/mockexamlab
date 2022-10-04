import React, { FC } from "react";
import QuestionContainer from "../../container/QuessionContainer";
import { useRouter } from "next/router";
import Instruction from "./Instruction";
import { RootStateOrAny, useSelector } from "react-redux";
import SpeakingQuestion from "./SpeakingQuestion";

const Speaking: FC = () => {
  const { test } = useSelector((state: RootStateOrAny) => state.tests);
  const { query } = useRouter();
  const speakingTest = test.sections.filter(
    (speak: any) => speak.componentType === "Speaking"
  );
  console.log(speakingTest);
  if (query.level == "intro") {
    return (
      <QuestionContainer title={""}>
        <Instruction
          title={speakingTest[0]?.componentType}
          time={"30"}
          name={""}
          section={"speaking"}
        />
      </QuestionContainer>
    );
  }
  return (
    <QuestionContainer title={speakingTest[0]?.componentType}>
      {speakingTest?.map((question: any, index: number) => (
        <SpeakingQuestion
          key={index}
          title={question.title}
          question={question.questionGroups[0]?.SpeakingTopic.topic}
          qNumber={question.questionGroups[0]?.SpeakingTopic.q_number}
        />
      ))}
    </QuestionContainer>
  );
};

export default Speaking;
