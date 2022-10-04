import React, { FC, useEffect } from "react";

import style from "styles/dashboard/exam.module.scss";
import { RootStateOrAny, useDispatch, useSelector } from "react-redux";
import { createTest, getTest } from "../../../app/ActionCreators";
import "react-toastify/dist/ReactToastify.css";
import { toast, ToastContainer } from "react-toastify";

interface Props {
  value: string;
}

const ExamResults: FC<Props> = ({ value }) => {
  const { tests } = useSelector((state: RootStateOrAny) => state.tests);
  const [results, setResults] = React.useState<
    { name: string; id: string; module: string; description: string }[]
  >([]);
  const dispatch = useDispatch();
  const createTestHandler = (id: string) => {
    if (id.length < 1) {
      return toast.error("Please select a exam");
    }
    dispatch(getTest(id));
  };
  useEffect(() => {
    if (value === "") return;
    // eslint-disable-next-line react-hooks/exhaustive-deps
    setResults(
      tests.filter((item: any) => item.name.toLowerCase().includes(value))
    );
  }, [tests, value]);
  if (!tests) return <h4>exams dose not exist.</h4>;
  return (
    <>
      <ToastContainer theme={"colored"} />
      <div className={style.exams__results}>
        {value.length < 1
          ? tests?.data?.map(
              (item: {
                name: string;
                id: string;
                module: string;
                description: string;
              }) => (
                <div
                  onClick={() => createTestHandler(item.id)}
                  className={style.exam}
                  key={item.id}
                >
                  <h4>{item.name.toUpperCase()}</h4>
                  <span>level: {item.module}</span>
                </div>
              )
            )
          : results.map((item, index) => (
              <div className={style.exam} key={index}>
                <h4>{item.name.toUpperCase()} </h4>
                <span>level: {item.module}</span>
              </div>
            ))}
      </div>
    </>
  );
};

export default ExamResults;
