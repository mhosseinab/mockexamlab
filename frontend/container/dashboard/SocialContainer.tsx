import DashboardLayout from "components/dashboard/DashboardLayout";
import Search from "components/dashboard/exam/Search";
import Header from "components/dashboard/shop/Header";
import Post from "components/dashboard/social/Post";
import DashboardHeader from "components/ui/DashboardHeader";
import Image from "next/image";
import Link from "next/link";
import React from "react";

import style from "styles/dashboard/social.module.scss";

const SocialContainer = () => {
  const [search, setSearch] = React.useState<string>("");
  return (
    <DashboardLayout>
      <DashboardHeader title="Stacks" />
      <Header
        height="226"
        image="/img/dashboard/school.png"
        link="/"
        title="Ask"
        width="226"
        br="Anything"
      />
      <Search setValue={setSearch} value={search} />
      <div className={style.social__container}>
        <h1>Latest posts</h1>
        <div className={style.posts}>
          {posts.map((post, index) => (
            <Post
              description={post.description}
              image={post.image}
              link={`/dashboard/social/post/${post.link}`}
              title={post.title}
              key={index}
            />
          ))}
        </div>
      </div>
    </DashboardLayout>
  );
};

const posts = [
  {
    image: "/img/dashboard/profile-pic.png",
    title: "Last and first name",
    description: "Hi friend! How do you do? Lets talk!",
    link: "post-1",
  },
  {
    image: "/img/dashboard/profile-pic.png",
    title: "Last and first name",
    description: "Hi friend! How do you do? Lets talk!",
    link: "post-2",
  },
  {
    image: "/img/dashboard/profile-pic.png",
    title: "Last and first name",
    description: "Hi friend! How do you do? Lets talk!",
    link: "post-3",
  },
  {
    image: "/img/dashboard/profile-pic.png",
    title: "Last and first name",
    description: "Hi friend! How do you do? Lets talk!",
    link: "post-4",
  },
];

export default SocialContainer;
