import Image from "next/image";
import React, { FC } from "react";

import style from "styles/dashboard/exam.module.scss";

interface Props {
  value: string;
  setValue: (event: string) => void;
}

const Search: FC<Props> = ({ setValue, value }) => {
  return (
    <div className={style.search}>
      <input
        type="text"
        value={value}
        onChange={(event) => setValue(event.target.value)}
        placeholder="Search any exams"
      />
      <span className={style.icon}>
        <Image
          src="/img/icons/search.svg"
          alt="search icon"
          width={21}
          height={21}
        />
      </span>
    </div>
  );
};

export default Search;
