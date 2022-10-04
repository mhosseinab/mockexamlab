import React from "react";
import style from "styles/exams/questions.module.scss";
import { Checkbox } from "@mui/material";

interface Props {
  title: string;
  question: string[];
}

const CheckQuestion: React.FC<Props> = ({ question, title }) => {
  const [selectedQuestion, setSelectedQuestion] = React.useState<string>();

  return (
    <div className={style.question__checkbox}>
      <p className={style.title}>{title}</p>
      {Object.values(question).map((item, index) => (
        <div
          className={style.question}
          key={index}
          onClick={() => setSelectedQuestion(item)}
        >
          <Checkbox
            id={item + index.toString()}
            value={item}
            checked={item === selectedQuestion}
            onClick={() => setSelectedQuestion(item)}
          />
          <label htmlFor={item}> {item} </label>
          <br />
        </div>
      ))}
    </div>
  );
};
export default CheckQuestion;
