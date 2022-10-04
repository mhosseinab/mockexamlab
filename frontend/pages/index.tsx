import HomeLayout from "components/ui/HomeLayout";
import Navbar from "components/ui/Navbar";
import HomeLanding from "container/HomeLanding";
import type { NextPage } from "next";

const Home: NextPage = () => {
  return (
    <>
      <HomeLayout>
        <Navbar links={[{ link: "/signup", title: "Sign up" },{ link: "/signin", title: "Sign in" }]} />
      </HomeLayout>
      <HomeLanding />
    </>
  );
};

export default Home;
