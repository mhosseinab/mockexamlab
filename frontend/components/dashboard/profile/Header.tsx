import Link from "next/link";
import React, {FC} from "react";

import style from "styles/dashboard/profile.module.scss";

interface Props {
    tags: {
        title: string;
        color?: string;
    }[];
}

const Header: FC<Props> = ({tags}) => {
    return (
        <div className={style.profile__header}>
            <div className={style.header}>
                <div className={style.banner}/>
                <div
                    className={style.profile__img}
                    style={{backgroundImage: "url('/img/dashboard/profile-pic.png')"}}
                />
            </div>
            <div className={style.details}>
                <div className={style.tags}>
                    {tags.map((tag, index) => (
                        <span
                            key={index}
                            style={{
                                backgroundColor: `${tag.color ? tag.color : "#3ac922"}`,
                            }}
                        >
              {tag.title}
            </span>
                    ))}
                </div>

                <Link href="/dashboard/financial">
                    <a>
                        wallet
                        <svg
                            width="12"
                            height="8"
                            viewBox="0 0 12 8"
                            fill="none"
                            xmlns="http://www.w3.org/2000/svg"
                        >
                            <path
                                fillRule="evenodd"
                                clipRule="evenodd"
                                d="M8.59942 6.70034C8.34475 7.04545 8.39652 7.54889 8.71506 7.82481C9.0336 8.10072 9.49827 8.04463 9.75295 7.69952L11.6228 5.16571C12.1257 4.4841 12.1257 3.5159 11.6228 2.83429L9.75295 0.300476C9.49827 -0.0446331 9.0336 -0.100725 8.71506 0.175193C8.39652 0.451109 8.34475 0.95455 8.59942 1.29966L10.0017 3.19996L0.738436 3.19996C0.33061 3.19996 0 3.55815 0 4C0 4.44185 0.33061 4.80004 0.738436 4.80004L10.0017 4.80004L8.59942 6.70034Z"
                                fill="white"
                            />
                        </svg>
                    </a>
                </Link>

            </div>
        </div>
    );
};

export default Header;
