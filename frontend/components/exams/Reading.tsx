import React from "react";
import Instruction from "./Instruction";
import QuestionContainer from "../../container/QuessionContainer";
import { RootStateOrAny, useSelector } from "react-redux";
import ReadingQuestion from "./ReadingQuesstion";
import { useRouter } from "next/router";

const Reading = () => {
  const { auth, tests } = useSelector((state: RootStateOrAny) => state);
  const { query } = useRouter();
  return (
    <QuestionContainer title={auth.userName}>
      {query.level == "intro" ? (
        <Instruction
          title={"Instruction"}
          time={"30"}
          name={"IELTS Reading"}
          section={"reading"}
        />
      ) : (
        <div>
          {tests.test.sections && (
            <>
              <ReadingQuestion />
            </>
          )}
        </div>
      )}
    </QuestionContainer>
  );
};
export default Reading;
