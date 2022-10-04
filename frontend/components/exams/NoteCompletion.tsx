import { FC } from "react";
import style from "styles/exams/questions.module.scss";
interface Props {
  question: string;
}

const NoteCompletion: FC<Props> = ({ question }) => {
  const html = question.replace(RegExp("{%.[1-9].%}.", "g"), `<input/>`);
  const findMatch = question.match(RegExp("{%.[1-9].%}.", "g"));
  return (
    <div className={style.note__completion}>
      <div
        dangerouslySetInnerHTML={{
          __html: html,
        }}
      />
    </div>
  );
};

export default NoteCompletion;
