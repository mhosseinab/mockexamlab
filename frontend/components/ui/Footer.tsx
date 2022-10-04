import React from "react";

import style from "/styles/home.module.scss"
import Image from "next/image";

const Footer = () => {
    return (
        <div className={style.footer}>
            <ul className={style.address}>
                <div className={style.logo}>
                    <Image src="/img/logo.png" alt="logo" layout="fill"/>
                </div>
                <li>
                    <h6>
                        Address
                    </h6>
                    <small> 4710-4890 breackinridge St. UK Burlington, VT 05401
                    </small>
                </li>
                <li>
                    <h6>
                        Need help?
                    </h6>
                    <small>
                        Call: 1-800-345-6789
                    </small>
                </li>
            </ul>
            <ul className={style.blog}>
                <h6>Latest blogs</h6>
                <li>
                    <Image src={"/img/blog.png"} width={76} height={80} alt={"blog"}/>
                    <div className={style.blog__details}>
                        <h6>3 how to grow marketing
                            ways</h6>
                        <p>by Mohsen Khani | 30 Oct, 2019</p>
                    </div>
                </li>
                <li>
                    <Image src={"/img/blog.png"} width={76} height={80} alt={"blog"}/>
                    <div className={style.blog__details}>
                        <h6>3 how to grow marketing
                            ways</h6>
                        <p>by Mohsen Khani | 30 Oct, 2019</p>
                    </div>
                </li>
            </ul>
            <ul className={style.information}>
                <h6>information</h6>
                <li>About us</li>
                <li>Contact us</li>
                <li>Terms and conditions</li>
                <li>Privacy policy</li>
                <li>Delivery information</li>
                <li>Brands</li>
            </ul>
        </div>
    );
};

export default Footer
