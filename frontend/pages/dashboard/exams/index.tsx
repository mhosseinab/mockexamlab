import ExamContainer from "container/ExamContainer";
import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { getAllExams } from "../../../app/ActionCreators";
import { getCookie } from "cookies-next";
import { RootState } from "../../../app/store";
import Spinner from "../../../components/ui/Spinner";

const Exam = () => {
  const [loading, setLoading] = React.useState(false);
  const dispatch = useDispatch();
  const { tests } = useSelector((state: RootState) => state.tests);
  useEffect(() => {
    setLoading(true);
    dispatch(getAllExams(getCookie("token") as string));
    if (tests) setLoading(false);
    setTimeout(() => {
      setLoading(false);
    }, 5000);
  }, []);
  return !loading ? <ExamContainer /> : <Spinner />;
};

export default Exam;
