import Image from "next/image";
import Link from "next/link";
import React, { FC } from "react";

import style from "styles/dashboard/profile.module.scss";

interface Props {
  title: string;
  link: string;
  items: { title: string; image: string; link?: string; price?: number }[];
}

const CardList: FC<Props> = ({ items, link, title }) => {
  return (
    <div className={style.card__list}>
      <div className={style.header}>
        <h1>{title}</h1>
        <Link href={link}>
          <a>View All</a>
        </Link>
      </div>
      <div className={style.cards}>
        {items.map((item, index) => (
          <div className={style.card} key={index}>
            <div className={style.img}>
              <Image src={item.image} alt={item.title} layout="fill" />
            </div>
            <div
              className={style.details}
              style={{
                display: "flex",
                justifyContent: "space-between",
                flexDirection: `${item.price ? "row" : "column"}`,
              }}
            >
              <h6>{item.title}</h6>
              {item.link && (
                <Link href={item.link}>
                  <a>More Detail</a>
                </Link>
              )}
              {item.price && (
                <h6 style={{ color: "#369926" }}>{item.price}$</h6>
              )}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default CardList;
