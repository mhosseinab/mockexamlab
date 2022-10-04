import DashboardHeader from "components/ui/DashboardHeader";
import Image from "next/image";
import Link from "next/link";
import React from "react";

import style from "/styles/dashboard/exam.module.scss";

const Header = () => {
    return (
        <div className={style.header}>
            <DashboardHeader title="Hello ,"/>
            <div className={style.cards}>
                <div className={style.card__one}>
                    <div className={style.icon}>
                        <Image
                            src="/img/icons/fire.svg"
                            alt="fire icon"
                            width={41}
                            height={52}
                        />
                    </div>
                    <div className={style.details}>
                        <h1>
                            Take your <br/> IELTS
                        </h1>
                        <Link href="/">
                            <a>
                                Today
                                <svg
                                    width="24"
                                    height="24"
                                    viewBox="0 0 24 24"
                                    fill="none"
                                    xmlns="http://www.w3.org/2000/svg"
                                >
                                    <path
                                        d="M5 12H19"
                                        stroke="#ffffff"
                                        strokeWidth={2}
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                    />
                                    <path
                                        d="M12 5L19 12L12 19"
                                        stroke="#ffffff"
                                        strokeWidth={2}
                                        strokeLinecap="round"
                                        strokeLinejoin="round"
                                    />
                                </svg>
                            </a>
                        </Link>
                    </div>
                </div>
                <div className={style.card__two}>
                    <div className={style.icon}>
                        <div className={style.img}>
                            <Image src="/img/dashboard/hat.png" alt="icon" layout="fill"/>
                        </div>
                    </div>
                    <div className={style.details}>
                        <h2>
                            Learn IELTS in <br/> 30 Days
                        </h2>
                        <div className={style.courses}>
                            <svg
                                width="12"
                                height="14"
                                viewBox="0 0 12 14"
                                fill="none"
                                xmlns="http://www.w3.org/2000/svg"
                            >
                                <path
                                    d="M1.99984 0.333328C1.2665 0.333328 0.673171 0.933328 0.673171 1.66666L0.666504 12.3333C0.666504 13.0667 1.25984 13.6667 1.99317 13.6667H9.99984C10.7332 13.6667 11.3332 13.0667 11.3332 12.3333V4.33333L7.33317 0.333328H1.99984ZM6.6665 5V1.33333L10.3332 5H6.6665Z"
                                    fill="white"
                                />
                            </svg>
                            15 course
                        </div>
                    </div>
                </div>
            </div>
            <Link href="/">
                <a className={style.check}>
                    Check all the routes
                    <svg
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                    >
                        <path
                            d="M5 12H19"
                            stroke="#ffffff"
                            strokeWidth={2}
                            strokeLinecap="round"
                            strokeLinejoin="round"
                        />
                        <path
                            d="M12 5L19 12L12 19"
                            stroke="#ffffff"
                            strokeWidth={2}
                            strokeLinecap="round"
                            strokeLinejoin="round"
                        />
                    </svg>
                </a>
            </Link>
        </div>
    );
};

export default Header;
