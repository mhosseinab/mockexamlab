import DashboardLayout from "components/dashboard/DashboardLayout";
import Post from "components/dashboard/social/Post";
import ChatCard from "components/ui/ChatCard";
import DashboardHeader from "components/ui/DashboardHeader";
import React from "react";
import style from "styles/dashboard/social.module.scss";

const PostView = () => {
  return (
    <DashboardLayout>
      <DashboardHeader title="A Post" />
      <ChatCard
        description="Amet minim mollit non deserunt ullamco est sit aliqua dolor do amet sint. Velit officia consequat duis enim velit mollit. Exercitation veniam consequat sunt."
        image="/img/dashboard/profile-pic.png"
      />
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

export default PostView;
