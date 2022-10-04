import React, { FC } from "react";

import style from "styles/dashboard/stats.module.scss";

interface Props {
  yourTime: string;
  otherTime: string;
}

const ExamTime: FC<Props> = ({ otherTime, yourTime }) => {
  return (
    <div className={style.timing}>
      <div className={style.time1}>
        <div className={style.your__time}>
          <div className={style.percentage} style={{ height: +yourTime }} />
          <h6>You</h6>
        </div>
      </div>
      <div className={style.time2}>
        <div className={style.other__time}>
          <div className={style.percentage} style={{ height: +otherTime }} />
          <h6>Others</h6>
        </div>
      </div>
    </div>
  );
};

export default ExamTime;
