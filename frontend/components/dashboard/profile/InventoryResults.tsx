import React, { FC } from "react";

import style from "styles/dashboard/profile.module.scss";

interface Props {
  currentMediaType: string;
  setCurrentMediaType: (mediaType: string) => void;
}

const InventoryResults: FC<Props> = ({
  currentMediaType,
  setCurrentMediaType,
}) => {
  return (
    <div className={style.inventory__results}>
      <div className={style.media__types}>
        {mediaTypes.map((mediaType, index) => (
          <span
            className={`${style.media__type} ${
              currentMediaType === mediaType && style.selected
            }`}
            onClick={() => setCurrentMediaType(mediaType)}
            key={index}
          >
            {mediaType}
          </span>
        ))}
      </div>
    </div>
  );
};

const mediaTypes = ["Event", "Books", "Videos"];

export default InventoryResults;
