import React, { useEffect } from "react";
import { RootStateOrAny, useDispatch, useSelector } from "react-redux";
import { getTest } from "app/ActionCreators";
import { useRouter } from "next/router";
import Reading from "../../../components/exams/Reading";
import { getCookie } from "cookies-next";
import Spinner from "../../../components/ui/Spinner";
import Listening from "../../../components/exams/Listening";
import Speaking from "../../../components/exams/Speaking";
import Writing from "../../../components/exams/Writing";

const Exam = () => {
  const { query } = useRouter();
  const dispatch = useDispatch();
  const { createdTest, test } = useSelector(
    (state: RootStateOrAny) => state.tests
  );
  useEffect(() => {
    if (!test.id) {
      dispatch(getTest(createdTest.testId || getCookie("testId")));
    }
  }, [dispatch, query.level, test.id, createdTest.testId]);
  if (!test.id) {
    return <Spinner />;
  }
  switch (query?.section?.toString().toLowerCase()) {
    case "reading":
      return <Reading />;
    case "listening":
      return <Listening />;
    case "speaking":
      return <Speaking />;
    case "writing":
      return <Writing />;
    default:
      return <Spinner />;
  }
};

export default Exam;
