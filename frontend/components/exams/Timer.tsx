import React, { FC } from "react";
import { useState, useEffect } from "react";
import style from "styles/exams/questions.module.scss";
import Countdown from "react-countdown";

interface Props {
  deadline: Date;
}

const Timer: FC<Props> = ({ deadline }) => {
  const renderer = ({ hours, minutes, seconds, completed }: any) => {
    if (completed) {
      // Render a completed state
      return <span>Time over</span>;
    } else {
      // Render a countdown
      return (
        <span>
          {minutes}:{seconds}
        </span>
      );
    }
  };
  return <Countdown date={deadline} renderer={renderer} />;
};

export default Timer;
