import Image from "next/image";
import Link from "next/link";
import React, {FC} from "react";

import style from "styles/dashboard/shop.module.scss";

interface Props {
    title: string;
    link: string;
    br?: string;
    image: string;
    width: string;
    height: string;
}

const Header: FC<Props> = ({image, link, title, br, height, width}) => {
    return (
        <div className={style.shop__header}>
            <div className={style.img}>
                <Image src={image} alt="trophy" width={width} height={height}/>
            </div>
            <div className={style.details}>
                <h2>
                    {title}
                    <br/> {br && br}
                </h2>
                <Link href={link}>
                    <a>
                        More Details
                        <svg
                            width="24"
                            height="24"
                            viewBox="0 0 24 24"
                            fill="none"
                            xmlns="http://www.w3.org/2000/svg"
                        >
                            <path
                                d="M5 12H19"
                                stroke="#fff"
                                strokeWidth={2}
                                strokeLinecap="round"
                                strokeLinejoin="round"
                            />
                            <path
                                d="M12 5L19 12L12 19"
                                stroke="#fff"
                                strokeWidth={2}
                                strokeLinecap="round"
                                strokeLinejoin="round"
                            />
                        </svg>
                    </a>
                </Link>
            </div>
        </div>
    );
};

export default Header;
