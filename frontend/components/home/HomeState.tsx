import Image from "next/image";
import React from "react";

import style from "/styles/home.module.scss";

const HomeState = () => {
  return (
    <div className={style.home__stats}>
      <div className={style.details}>
        <h1>Stats!</h1>
        <p>
          We are not leaving you! infact, we monitor your progress and give you <br/>
          hints to learning in a better way
        </p>
      </div>
      <div className={style.chart}>
        <Image
          src="/img/chart-blue.png"
          alt="chart-one"
          width={246}
          height={156}
        />
        <span>Your grades</span>
      </div>
      <div className={style.chart}>
        <Image
          src="/img/chart-yellow.png"
          alt="chart-one"
          width={246}
          height={156}
        />
        <span>Your courses</span>
      </div>
    </div>
  );
};

export default HomeState;
