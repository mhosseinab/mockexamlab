import Image from "next/image";
import Link from "next/link";
import React, {FC, useEffect} from "react";
import style from "styles/menu.module.scss";
import {getCookie} from "cookies-next"

interface Props {
    links: {
        title: string;
        link: string;
    }[];
}

const Navbar: FC<Props> = ({links}) => {
    const allLinks = getCookie("token") ? [{link: "/dashboard", title: "Dashboard"}] : links
    const [isSSR, setIsSSR] = React.useState(true);

    useEffect(() => {
        setIsSSR(false);
    }, []);
    return (
        <nav className={style.navbar}>
            <div className={style.logo}>
                <Image src="/img/logo.png" alt="logo" layout="fill"/>
            </div>
            <ul className={style.menu}>
                <li>
                    <Link href="/exams">
                        <a>
                            <Image
                                src="/img/icons/save.png"
                                alt="icon"
                                width={16}
                                height={18}
                            />
                            Exams
                        </a>
                    </Link>
                </li>
                <li>
                    <Link href="/courses">
                        <a>
                            <Image
                                src="/img/icons/video.png"
                                alt="icon"
                                width={16}
                                height={18}
                            />
                            Courses
                        </a>
                    </Link>
                </li>
                <li>
                    <Link href="/shop">
                        <a>
                            <Image
                                src="/img/icons/shop.png"
                                alt="icon"
                                width={16}
                                height={18}
                            />
                            Shop
                        </a>
                    </Link>
                </li>
            </ul>
            {
                !isSSR && (
                    <div className={style.auth__btn}>
                        {allLinks.map((link, index) => (
                            <Link key={index} href={link.link}>
                                <a className={style.btn}>{link.title}</a>
                            </Link>
                        ))}
                    </div>
                )
            }
        </nav>
    );
};

export default Navbar;
