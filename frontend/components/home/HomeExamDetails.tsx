import React from "react";

import style from "styles/home.module.scss";

const HomeExamDetails = () => {
  const [currentExam, setCurrentExam] = React.useState(exams[0]);
  return (
    <div className={style.home__examDetails}>
      <div className={style.mobile__title}>
        <h3>Why using us?</h3>
        <h4>Fast and online</h4>
      </div>
      <div className={style.exams}>
        {exams.map((item, index) => (
          <div
            key={index}
            className={`${style.exam} ${
              currentExam.id === item.id && style.active
            }`}
            onClick={() => setCurrentExam(item)}
          >
            <svg
              width="47"
              height="48"
              viewBox="0 0 47 48"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M21.211 0.666687C18.6337 0.666687 16.5443 2.75602 16.5443 5.33335V7.55826L16.3382 7.66131L14.6771 6.00018C12.8547 4.17773 9.89988 4.17773 8.07743 6.00018L5.37695 8.70066C3.5545 10.5231 3.5545 13.4779 5.37695 15.3003L7.03808 16.9615L6.93503 17.1676L4.71012 17.1676C2.1328 17.1676 0.043457 19.2569 0.043457 21.8342V26.168C0.043457 28.7454 2.1328 30.8347 4.71012 30.8347H6.93503L7.03808 31.0408L5.37695 32.7019C3.5545 34.5244 3.5545 37.4791 5.37695 39.3016L8.07743 42.0021C9.89988 43.8245 12.8547 43.8245 14.6771 42.0021L16.3382 40.3409L16.5443 40.444L16.5443 42.6689C16.5443 45.2462 18.6337 47.3356 21.211 47.3356H25.5448C28.1221 47.3356 30.2115 45.2462 30.2115 42.6689V40.444L30.4176 40.3409L32.0787 42.0021C33.9011 43.8245 36.8559 43.8245 38.6784 42.0021L41.3788 39.3016C43.2013 37.4791 43.2013 34.5244 41.3788 32.7019L39.7177 31.0408L39.8208 30.8347H42.0457C44.623 30.8347 46.7123 28.7454 46.7123 26.168V21.8342C46.7123 19.2569 44.623 17.1676 42.0457 17.1676H39.8208L39.7177 16.9615L41.3788 15.3003C43.2013 13.4779 43.2013 10.5231 41.3788 8.70066L38.6784 6.00018C36.8559 4.17773 33.9011 4.17773 32.0787 6.00018L30.4176 7.66131L30.2115 7.55826V5.33335C30.2115 2.75602 28.1221 0.666687 25.5448 0.666687H21.211ZM21.211 5.33335H25.5448V7.55826C25.5448 9.32586 26.5435 10.9418 28.1245 11.7322L28.3306 11.8353C30.1272 12.7336 32.2971 12.3815 33.7174 10.9611L35.3785 9.30001L38.079 12.0005L36.4179 13.6616C34.9975 15.082 34.6454 17.2518 35.5437 19.0484L35.6468 19.2546C36.4373 20.8355 38.0532 21.8342 39.8208 21.8342H42.0457V26.168H39.8208C38.0532 26.168 36.4373 27.1667 35.6468 28.7477L35.5437 28.9538C34.6454 30.7504 34.9975 32.9203 36.4179 34.3406L38.079 36.0018L35.3785 38.7022L33.7174 37.0411C32.2971 35.6208 30.1272 35.2686 28.3306 36.1669L28.1245 36.27C26.5435 37.0605 25.5448 38.6764 25.5448 40.444V42.6689H21.211V40.444C21.211 38.6764 20.2123 37.0605 18.6313 36.27L18.4252 36.1669C16.6286 35.2686 14.4587 35.6208 13.0384 37.0411L11.3773 38.7022L8.67678 36.0018L10.3379 34.3406C11.7583 32.9203 12.1104 30.7504 11.2121 28.9538L11.109 28.7477C10.3185 27.1667 8.70263 26.168 6.93503 26.168H4.71012V21.8342H6.93503C8.70263 21.8342 10.3185 20.8355 11.109 19.2546L11.2121 19.0485C12.1104 17.2518 11.7583 15.082 10.3379 13.6616L8.67678 12.0005L11.3773 9.30001L13.0384 10.9611C14.4587 12.3815 16.6286 12.7336 18.4252 11.8353L18.6313 11.7322C20.2123 10.9418 21.211 9.32586 21.211 7.55826V5.33335ZM19.7109 24.0002C19.7109 21.9754 21.3523 20.3341 23.377 20.3341C25.4018 20.3341 27.0431 21.9754 27.0431 24.0002C27.0431 26.0249 25.4018 27.6663 23.377 27.6663C21.3523 27.6663 19.7109 26.0249 19.7109 24.0002ZM23.377 15.6674C18.775 15.6674 15.0442 19.3981 15.0442 24.0002C15.0442 28.6022 18.775 32.3329 23.377 32.3329C27.9791 32.3329 31.7098 28.6022 31.7098 24.0002C31.7098 19.3981 27.9791 15.6674 23.377 15.6674Z"
                fill="#52C3FF"
              />
            </svg>
            <h5>{item.title}</h5>
            <p>{item.description}</p>
          </div>
        ))}
      </div>
      <div className={style.details}>
        <p>{currentExam.details}</p>
      </div>
    </div>
  );
};

const exams = [
  {
    id: 1,
    title: "Exams",
    description: "over than 20 exams in one place!",
    details:
      "Li Europan lingues es membres del sam familie. Lor separat existentie es un myth. Por scientie, musica, sport etc., li tot Europa usa li sam vocabularium. Li lingues differe solmen in li grammatica, li pronunciation e li plu commun vocabules. Omnicos directe al desirabilit? de un nov lingua franca: on refusa continuar payar custosi traductores. It solmen va esser necessi far uniform grammatica, pronunciation e plu sommun paroles.",
  },
  {
    id: 2,
    title: "Exams",
    description: "over than 20 exams in one place!",
    details:
      "Li Europan lingues es membres del sam familie. Lor separat existentie es un myth. Por scientie, musica, sport etc., li tot Europa usa li sam vocabularium. Li lingues differe solmen in li grammatica, li pronunciation e li plu commun vocabules. Omnicos directe al desirabilit? de un nov lingua franca: on refusa continuar payar custosi traductores. It solmen va esser necessi far uniform grammatica, pronunciation e plu sommun paroles.",
  },
  {
    id: 3,
    title: "Exams",
    description: "over than 20 exams in one place!",
    details:
      "Li Europan lingues es membres del sam familie. Lor separat existentie es un myth. Por scientie, musica, sport etc., li tot Europa usa li sam vocabularium. Li lingues differe solmen in li grammatica, li pronunciation e li plu commun vocabules. Omnicos directe al desirabilit? de un nov lingua franca: on refusa continuar payar custosi traductores. It solmen va esser necessi far uniform grammatica, pronunciation e plu sommun paroles.",
  },
  {
    id: 4,
    title: "Exams",
    description: "over than 20 exams in one place!",
    details:
      "Li Europan lingues es membres del sam familie. Lor separat existentie es un myth. Por scientie, musica, sport etc., li tot Europa usa li sam vocabularium. Li lingues differe solmen in li grammatica, li pronunciation e li plu commun vocabules. Omnicos directe al desirabilit? de un nov lingua franca: on refusa continuar payar custosi traductores. It solmen va esser necessi far uniform grammatica, pronunciation e plu sommun paroles.",
  },
];

export default HomeExamDetails;
