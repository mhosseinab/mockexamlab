import DashboardLayout from "components/dashboard/DashboardLayout";
import ExamTime from "components/dashboard/stats/ExamTime";
import DashboardHeader from "components/ui/DashboardHeader";
import Image from "next/image";
import React from "react";

import style from "styles/dashboard/stats.module.scss";

const StatsContainer = () => {
  const [currentTab, setCurrentTab] = React.useState("overall");
  const [currentGraph, setCurrentGraph] = React.useState<{
    yourTime: string;
    otherTime: string;
    title: string;
  }>(exams[0]);
  return (
    <DashboardLayout>
      <DashboardHeader title="Shop invoice" />
      <div className={style.menu}>
        <button
          className={`${currentTab == "exam" && style.selected}`}
          onClick={() => setCurrentTab("exam")}
        >
          Per Exam
        </button>
        <button
          className={`${currentTab == "overall" && style.selected}`}
          onClick={() => setCurrentTab("overall")}
        >
          Overall
        </button>
      </div>
      <div className={style.chart}>
        <h3>Your {currentTab} grades</h3>
        <div className={style.img}>
          {currentTab == "exam" ? (
            <Image src="/img/dashboard/chart.png" alt="chart" layout="fill" />
          ) : (
            <Image src="/img/dashboard/chart.png" alt="chart" layout="fill" />
          )}
        </div>
      </div>
      <div className={style.times}>
        <div className={style.exams__title}>
          {exams.map((exam, index) => (
            <span
              onClick={() => setCurrentGraph(exam)}
              className={`${exam == currentGraph && style.selected} ${
                style.exam__title
              }`}
              key={index}
            >
              {exam.title.toUpperCase()}
            </span>
          ))}
        </div>
        <div className={style.graphs}>
          <ExamTime
            otherTime={currentGraph.otherTime}
            yourTime={currentGraph.yourTime}
          />

          <ExamTime
            otherTime={currentGraph.otherTime}
            yourTime={currentGraph.yourTime}
          />
        </div>
      </div>
    </DashboardLayout>
  );
};

const exams = [
  { title: "ietls", yourTime: "40", otherTime: "80" },
  { title: "toefl", yourTime: "50", otherTime: "70" },
  { title: "gbtd", yourTime: "30", otherTime: "40" },
];

export default StatsContainer;
