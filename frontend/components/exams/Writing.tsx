import React from "react";
import QuestionContainer from "../../container/QuessionContainer";
import Instruction from "./Instruction";
import { useRouter } from "next/router";
import WritingQuestion from "./WritingQuestion";
import { RootStateOrAny, useSelector } from "react-redux";

const Writing = () => {
  const { query } = useRouter();
  const { test } = useSelector((state: RootStateOrAny) => state.tests);
  const writingTest = test.sections.filter(
    (item: any) => item.componentType == "Writing"
  );
  if (query.level == "intro")
    return (
      <QuestionContainer title={""}>
        <Instruction
          title={"Writing"}
          time={"30"}
          name={"writing"}
          section={"writing"}
        />
      </QuestionContainer>
    );
  return (
    <QuestionContainer title={""}>
      <WritingQuestion question={writingTest[0]} />
    </QuestionContainer>
  );
};

export default Writing;
