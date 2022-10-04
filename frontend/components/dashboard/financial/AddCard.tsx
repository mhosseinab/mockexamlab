import React, { FC } from "react";

import style from "styles/dashboard/financial.module.scss";

interface Props {
  value: string;
  setValue: (value: string) => void;
}

const AddCard: FC<Props> = ({ setValue, value }) => {
  return (
    <div className={style.add__card}>
      <input
        type="text"
        placeholder="Text"
        value={value}
        onChange={(event) => setValue(event.target.value)}
      />
      <button>Add to cart</button>
    </div>
  );
};

export default AddCard;
