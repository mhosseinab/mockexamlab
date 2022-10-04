import React, { FC } from "react";
import style from "styles/exams/questions.module.scss";
import { useRouter } from "next/router";

interface Props {
  sections: {
    title: string;
    questionGroups: { length: number }[];
    link: string;
    componentType: string;
  }[];
}
const ExamsFooter: FC<Props> = ({ sections }) => {
  const { push, pathname, query } = useRouter();
  const currentSection = sections.filter(
    (sec) => sec.componentType.toLowerCase() == query.section
  );
  const changeTestHandler = () => {
    switch (query.section) {
      case "reading":
        return push({
          pathname,
          query: { section: "writing", level: "intro" },
        });
      case "writing":
        return push({
          pathname,
          query: { section: "listening", level: "intro" },
        });
      case "listening":
        return push({
          pathname,
          query: { section: "speaking", level: "intro" },
        });
      case "speaking":
        return push("/dashboard/exams");
    }
  };
  return (
    <div className={style.questions__footer}>
      <div className={style.forwards}>
        <div className={style.icon}>
          <svg
            width="10"
            height="21"
            viewBox="0 0 10 21"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              fillRule="evenodd"
              clipRule={"evenodd"}
              d="M9.65261 19.9583C10.0798 19.5914 10.1183 18.9593 9.73856 18.5466L2.06995 10.2109L9.73857 1.87529C10.1183 1.46251 10.0798 0.830436 9.65262 0.463518C9.2254 0.0966011 8.57122 0.133781 8.19146 0.546564L0.522853 8.8822C-0.174282 9.63997 -0.174282 10.7819 0.522852 11.5397L8.19146 19.8753C8.57121 20.2881 9.22539 20.3253 9.65261 19.9583Z"
              fill="white"
            />
          </svg>
        </div>
        <div className={style.icon}>
          <svg
            width="10"
            height="21"
            viewBox="0 0 10 21"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              fillRule="evenodd"
              clipRule="evenodd"
              d="M0.335701 19.9583C-0.0770818 19.5914 -0.114264 18.9593 0.252655 18.5466L7.66211 10.2109L0.252652 1.87529C-0.114267 1.46251 -0.0770848 0.830436 0.335698 0.463518C0.74848 0.096601 1.38055 0.133781 1.74747 0.546563L9.15693 8.8822C9.8305 9.63997 9.8305 10.7819 9.15693 11.5397L1.74747 19.8753C1.38056 20.2881 0.748484 20.3253 0.335701 19.9583Z"
              fill="white"
            />
          </svg>
        </div>
      </div>
      <div className={style.pagination}>
        {currentSection.map((sec, index) => (
          <div key={index}>
            <h4>{sec.title.charAt(0).toUpperCase() + sec.title.slice(1)}</h4>
            {sec.questionGroups.map((gp, index) => (
              <button className={style.paginate} key={index}>
                {index + 1}
              </button>
            ))}
          </div>
        ))}
      </div>
      <button onClick={changeTestHandler} className={style.finish}>
        Finish
      </button>
    </div>
  );
};

export default ExamsFooter;
