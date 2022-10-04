import React from "react";
import style from "/styles/exams/reading.module.scss";
import { useSelector } from "react-redux";
import { RootState } from "../../app/store";
import TrueOrFalseQuestion from "./TrueOrFasleQuestion";
import { log } from "util";

interface Props {}

const ReadingQuestion: React.FC<Props> = () => {
  const { test } = useSelector((state: RootState) => state.tests);
  return (
    <div className={style.reading__container}>
      <div className={style.reading}>
        <div className="sections__header">
          <h4>{test.sections[2].title}</h4>
        </div>
        <p>{test.sections[2].sectionReading.mainText}</p>
      </div>
      <div className={style.questions__container}>
        <div className={style.title}>
          <h4>{test.sections[2].componentType}</h4>
          <div className={style.br} />
        </div>
        <h4>
          Questions {"  "}
          {
            test.sections[2].questionGroups[0].identifyingInformation[0]
              .q_number
          }
          -{" "}
          {test.sections[2].questionGroups[0].identifyingInformation[0]
            .q_number +
            test.sections[2].questionGroups[0].identifyingInformation.length}
        </h4>
        <h4> {test.sections[2].questionGroups[0].description}</h4>
        <div className={style.questions}>
          {test.sections[2].questionGroups[0].identifyingInformation.map(
            (item: any, index: number) => (
              <TrueOrFalseQuestion
                key={index}
                title={item.title}
                count={item.q_number}
              />
            )
          )}
        </div>
      </div>
    </div>
  );
};
export default ReadingQuestion;
