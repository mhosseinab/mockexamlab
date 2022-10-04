import Link from "next/link";
import React, {FC} from "react";

import style from "styles/dashboard/shop.module.scss";

interface Props {
    category: string;
    setCategory: (category: string) => void;
}

const Routing: FC<Props> = ({category, setCategory}) => {
    return (
        <div className={style.shop__routing}>
            <div className={style.title}>
                <h1>Category</h1>
                <Link href="/">
                    <a>
                        See All
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
            <div className={style.routs}>
                {routs.map((route, index) => (
                    <Link href={`/dashboard/shop/${route.title.toLowerCase()}`} key={index}>
                        <div
                            className={`${style.route} ${
                                route.title === category && style.selected
                            } `}
                            onClick={() => setCategory(route.title)}
                        >
                            <span>{route.icon}</span>
                            <h6>{route.title}</h6>
                        </div>
                    </Link>
                ))}
            </div>
        </div>
    );
};

const routs = [
    {
        title: "Books",
        icon: (
            <svg
                width="12"
                height="14"
                viewBox="0 0 12 14"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
            >
                <path
                    fillRule="evenodd"
                    clipRule="evenodd"
                    d="M1.56522 1.90873C1.56522 1.75814 1.68201 1.63606 1.82609 1.63606H10.1739C10.318 1.63606 10.4348 1.75814 10.4348 1.90873V12.0886C10.4348 12.2945 10.2249 12.4261 10.0512 12.3292L6.85933 10.5499C6.32225 10.2505 5.67775 10.2505 5.14067 10.5499L1.94885 12.3292C1.77508 12.4261 1.56522 12.2945 1.56522 12.0886V1.90873ZM1.82609 0C0.817567 0 0 0.854568 0 1.90873V12.0886C0 13.5296 1.46905 14.4509 2.68542 13.7728L5.87724 11.9935C5.95396 11.9507 6.04604 11.9507 6.12276 11.9935L9.31458 13.7728C10.5309 14.4509 12 13.5296 12 12.0886V1.90873C12 0.854568 11.1824 0 10.1739 0H1.82609ZM3.91304 6.54422C3.48082 6.54422 3.13043 6.91047 3.13043 7.36225C3.13043 7.81403 3.48082 8.18028 3.91304 8.18028H8.08696C8.51918 8.18028 8.86957 7.81403 8.86957 7.36225C8.86957 6.91047 8.51918 6.54422 8.08696 6.54422H3.91304Z"
                    fill="#ffffff"
                />
            </svg>
        ),
    },
    {
        title: "Videos",
        icon: (
            <svg
                width="14"
                height="10"
                viewBox="0 0 14 10"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
            >
                <path
                    fillRule="evenodd"
                    clipRule="evenodd"
                    d="M0.777766 0C0.348218 0 0 0.353459 0 0.789474V9.21053C0 9.64654 0.348218 10 0.777766 10H9.76046C10.19 10 10.5382 9.64654 10.5382 9.21053V7.65855L12.7308 9.473C12.9635 9.6656 13.2847 9.705 13.5559 9.57424C13.8271 9.44347 14 9.1659 14 8.86111L14 1.06246C14 0.759093 13.8287 0.482558 13.5594 0.351037C13.2901 0.219517 12.9704 0.256277 12.7368 0.445608L10.5382 2.22817V0.789474C10.5382 0.353459 10.19 0 9.76046 0H0.777766ZM10.5382 5.62136V4.24897L12.4445 2.70345L12.4445 7.19887L10.5382 5.62136ZM8.9827 3.77945V2.89474V1.57895H1.55553V8.42105H8.9827V6.08589C8.97902 6.05483 8.97715 6.02339 8.97715 5.99173L8.97714 3.87366C8.97714 3.84198 8.97902 3.81052 8.9827 3.77945Z"
                    fill="white"
                />
            </svg>
        ),
    },
    {
        title: "Package",
        icon: (
            <svg
                width="12"
                height="16"
                viewBox="0 0 12 16"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
            >
                <path
                    d="M10.5 0.5H1.5C0.675 0.5 0 1.175 0 2V14C0 14.825 0.675 15.5 1.5 15.5H10.5C11.325 15.5 12 14.825 12 14V2C12 1.175 11.325 0.5 10.5 0.5ZM1.5 2H5.25V8L3.375 6.875L1.5 8V2Z"
                    fill="white"
                />
            </svg>
        ),
    },
];

export default Routing;
