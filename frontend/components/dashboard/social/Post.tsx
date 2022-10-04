import Image from "next/image";
import Link from "next/link";
import React, { FC } from "react";
import style from "styles/dashboard/social.module.scss";

interface Props {
  image: string;
  title: string;
  description: string;
  link: string;
}

const Post: FC<Props> = ({ description, image, link, title }) => {
  return (
    <Link href={link}>
      <div className={style.post__container}>
        <div className={style.profile}>
          <Image src={image} alt={title} layout="fill" />
          <div className={style.status} />
        </div>
        <div className={style.details}>
          <h4>{title}</h4>
          <p>{description}</p>
        </div>
      </div>
    </Link>
  );
};

export default Post;
