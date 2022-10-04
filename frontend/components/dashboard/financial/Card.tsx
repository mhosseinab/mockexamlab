import Image from "next/image";
import Link from "next/link";
import React from "react";

import style from "styles/dashboard/financial.module.scss";

const Card = () => {
  return (
    <div className={style.card}>
      <div className={style.head}>
        <div className="details">
          <Link href="/dashboard/financial">
            <h6>My Wallet</h6>
          </Link>
          <span>Balance</span>
          <h2>$ 430.90</h2>
        </div>
        <div className={style.logo}>
          <Image src="/img/icons/right.svg" alt="" layout="fill" />
          <Image src="/img/icons/left.svg" alt="" layout="fill" />
        </div>
      </div>
      <div className={style.bottom}>
        <Image
          src="/img/icons/dollar.svg"
          alt=""
          height={16.67}
          width={16.67}
        />
        <span>Deposit</span>
        <Image
          src="/img/icons/withdraw.svg"
          alt=""
          height={11.67}
          width={16.67}
        />
        <span>Withdraw</span>
      </div>
    </div>
  );
};

export default Card;
