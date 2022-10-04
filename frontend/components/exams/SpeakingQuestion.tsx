import React, { FC } from "react";
import style from "styles/exams/questions.module.scss";
import AudioRecorder from "./AudioRecorder";

type Props = {
  question: string;
  qNumber?: string;
  audio?: string;
  title: string;
};

const SpeakingQuestion: FC<Props> = ({ question, title, qNumber, audio }) => {
  return (
    <div className={style.speaking__question}>
      <div className={style.question}>
        <div className="sections__header">
          <h4>Speaking</h4>
          <p>{title}</p>
        </div>
        {title == "part 1" ? (
          <p>
            Question {qNumber} : {question}
          </p>
        ) : (
          <div dangerouslySetInnerHTML={{ __html: question }}></div>
        )}
        {audio && (
          <audio controls>
            <source src={audio} type="audio/mpeg" />
          </audio>
        )}
      </div>
      <div className={style.recorder}>
        <AudioRecorder />
      </div>
    </div>
  );
};

export default SpeakingQuestion;
