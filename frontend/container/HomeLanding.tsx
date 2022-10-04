import HomeExam from "components/home/HomeExam";
import HomeExamDetails from "components/home/HomeExamDetails";
import HomeHeader from "components/home/HomeHeader";
import HomeState from "components/home/HomeState";
import HomeWorks from "components/home/HomeWorks";
import React from "react";
import Footer from "../components/ui/Footer";

const HomeLanding = () => {
    return (
        <>
            <HomeHeader/>
            <HomeExam/>
            <HomeExamDetails/>
            <HomeState/>
            <HomeWorks/>
            <Footer/>
        </>
    );
};

export default HomeLanding;
