import React from "react";

import style from "styles/dashboard/financial.module.scss";

const Payments = () => {
    return (
        <div className={style.financial__payments}>
            {payments.map((payment, index) => (
                <div className={style.payment} key={index}>
                    <div className={style.title}>
                        <svg
                            width="20"
                            height="20"
                            viewBox="0 0 20 20"
                            fill="none"
                            xmlns="http://www.w3.org/2000/svg"
                        >
                            <path
                                fillRule="evenodd"
                                clipRule="evenodd"
                                d="M2 10C2 5.58172 5.58172 2 10 2C14.4183 2 18 5.58172 18 10C18 14.4183 14.4183 18 10 18C5.58172 18 2 14.4183 2 10ZM10 0C4.47715 0 0 4.47715 0 10C0 15.5228 4.47715 20 10 20C15.5228 20 20 15.5228 20 10C20 4.47715 15.5228 0 10 0ZM10 4.49969C10.5523 4.49969 11 4.94741 11 5.49969V11.4678C11 12.02 10.5523 12.4678 10 12.4678C9.44771 12.4678 9 12.02 9 11.4678V5.49969C9 4.94741 9.44771 4.49969 10 4.49969ZM11 14.501C11 15.0529 10.5523 15.5003 10 15.5003C9.44771 15.5003 9 15.0529 9 14.501C9 13.9491 9.44771 13.5017 10 13.5017C10.5523 13.5017 11 13.9491 11 14.501Z"
                                fill="#7C7B7B"
                            />
                        </svg>
                        <h3>{payment.title}</h3>
                    </div>
                    <span>{payment.date}</span>
                </div>
            ))}
        </div>
    );
};

const payments = [
    {title: "Payment 1", date: "01/01/2020"},
    {title: "Payment 2", date: "01/01/2020"},
    {title: "Payment 3", date: "01/01/2020"},
    {title: "Payment 4", date: "01/01/2020"},
];

export default Payments;
