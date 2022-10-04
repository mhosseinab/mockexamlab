import React, { FC } from "react";
import style from "styles/exams/questions.module.scss";
import Image from "next/image";
import Timer from "./Timer";
import { useRouter } from "next/router";
import AudioPlayer from "./AudioPlayer";

interface Props {
  title: string;
}

const TestHeader: FC<Props> = ({ title }) => {
  const [isPlaying, setISplaying] = React.useState(true);
  const THREE_DAYS_IN_MS = 1000 * 60 * 30;
  const NOW_IN_MS = new Date().getTime();
  const { query } = useRouter();
  return (
    <div className={style.question__layout}>
      <div className={style.header}>
        {query.section == "listening" && query.level != "intro" && (
          <AudioPlayer isPlaying={isPlaying} />
        )}
        <div className={style.test_title}>
          <div className={style.icon}>
            <Image
              src="/img/icons/person.svg"
              alt="person icon"
              width={14}
              height={20}
            />
          </div>
          <h3>{title}</h3>
        </div>
        <div className={style.exam__timer}>
          <div className={style.bar} />
          <div className={style.exam__time}>
            {query.level == "intro" ? (
              <>30:00</>
            ) : (
              <Timer deadline={new Date(THREE_DAYS_IN_MS + NOW_IN_MS)} />
            )}
            <svg
              width="14"
              height="14"
              viewBox="0 0 14 14"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M7.00033 13.8337C10.7743 13.8337 13.8337 10.7743 13.8337 7.00033C13.8337 3.22638 10.7743 0.166992 7.00033 0.166992C3.22638 0.166992 0.166992 3.22638 0.166992 7.00033C0.166992 10.7743 3.22638 13.8337 7.00033 13.8337ZM12.3337 7.00033C12.3337 9.94584 9.94584 12.3337 7.00033 12.3337C4.05481 12.3337 1.66699 9.94584 1.66699 7.00033C1.66699 4.05481 4.05481 1.66699 7.00033 1.66699C9.94584 1.66699 12.3337 4.05481 12.3337 7.00033ZM7.75033 3.66699C7.75033 3.25278 7.41454 2.91699 7.00033 2.91699C6.58611 2.91699 6.25033 3.25278 6.25033 3.66699V6.5861C6.25033 7.05022 6.4347 7.49534 6.76288 7.82353L7.80332 8.86399C8.09622 9.15688 8.57109 9.15688 8.86399 8.86399C9.15688 8.5711 9.15688 8.09623 8.86399 7.80333L7.82355 6.76287C7.77666 6.71599 7.75033 6.6524 7.75033 6.5861V3.66699Z"
                fill="#52C3FF"
              />
            </svg>
          </div>
        </div>
        <div className={style.sound}>
          <div className={style.bar} />
          <div className={style.actions}>
            <svg
              width="18"
              height="13"
              viewBox="0 0 18 13"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M10.6645 1.79447C10.666 0.376838 8.93044 -0.36085 7.86336 0.603863L5.30168 2.91979L2.32721 2.93125C1.4081 2.93479 0.665381 3.66193 0.666995 4.55666L0.670921 6.73378V8.242C0.670921 9.13063 1.40521 9.85395 2.318 9.86447L5.29686 9.8988L7.84833 12.2254C8.91244 13.1957 10.6531 12.4622 10.6546 11.0429L10.6645 1.79447ZM6.43596 4.10872L8.99764 1.79279L8.98783 11.0412L6.43636 8.71461C6.13245 8.43749 5.73299 8.28114 5.3166 8.27634L2.33774 8.242L2.33774 6.73305L2.33381 4.55382L5.30828 4.54235C5.72677 4.54074 6.12932 4.38595 6.43596 4.10872ZM16.212 3.46268C15.9162 3.11939 15.3905 3.07453 15.0379 3.36248C14.6861 3.64968 14.6394 4.1595 14.9327 4.50277L14.935 4.50555C14.9382 4.50945 14.9445 4.51713 14.9534 4.52837C14.9618 4.53895 14.9724 4.55264 14.985 4.56926C14.9992 4.58801 15.0159 4.61049 15.0345 4.63647C15.105 4.73505 15.2013 4.8807 15.2976 5.06086C15.4951 5.43026 15.6668 5.89573 15.6668 6.37503C15.6668 6.83486 15.4779 7.31317 15.2477 7.71275C15.1367 7.90552 15.0253 8.06398 14.9428 8.17284C14.9018 8.22693 14.8686 8.26784 14.8471 8.29364C14.8363 8.30651 14.8286 8.31554 14.8243 8.32049L14.8214 8.32385L14.8206 8.3247C14.5191 8.66164 14.5542 9.17277 14.8996 9.46757C15.2457 9.76296 15.7723 9.7293 16.0757 9.39239L15.4491 8.85755C16.0757 9.39239 16.0762 9.39185 16.0764 9.39167L16.0771 9.39084L16.0789 9.38884L16.0836 9.38354L16.0974 9.36775C16.1085 9.35496 16.1234 9.33763 16.1414 9.31602C16.1774 9.27284 16.2263 9.2123 16.2838 9.13642C16.3984 8.98533 16.5498 8.7698 16.7016 8.50636C16.997 7.99352 17.3337 7.23058 17.3337 6.37503C17.3337 5.53895 17.042 4.80893 16.7761 4.3117C16.6407 4.05853 16.5053 3.85307 16.4021 3.70887C16.3503 3.63649 16.306 3.57875 16.2732 3.53731C16.2567 3.51658 16.2431 3.49986 16.2328 3.48737L16.2198 3.4718L16.2152 3.46643L16.2134 3.46435L16.212 3.46268ZM15.5735 3.98406C16.212 3.46268 16.2118 3.46249 16.212 3.46268L15.5735 3.98406Z"
                fill="#7C7B7B"
              />
            </svg>
            <svg
              width="18"
              height="13"
              viewBox="0 0 18 13"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M10.8679 1.79447C10.8694 0.376838 9.132 -0.36085 8.06373 0.603863L5.49921 2.91979L2.52144 2.93125C1.6013 2.93479 0.857762 3.66193 0.859378 4.55666L0.863308 6.73305V8.242C0.863308 9.13063 1.59841 9.85395 2.51222 9.86447L5.49438 9.8988L8.04869 12.2254C9.11398 13.1957 10.8566 12.4622 10.8581 11.0429L10.8679 1.79447ZM6.63475 4.10872L9.19927 1.79279L9.18945 11.0412L6.63515 8.71461C6.3309 8.43749 5.93099 8.28114 5.51414 8.27634L2.53197 8.242L2.53198 6.73305L2.53197 6.73162L2.52804 4.55382L5.50582 4.54235C5.92477 4.54074 6.32776 4.38595 6.63475 4.10872ZM14.6445 1.64536C15.0041 1.36523 15.5292 1.42161 15.8173 1.77131L15.1661 2.27853C15.8173 1.77131 15.8173 1.77131 15.8173 1.77131L15.818 1.77214L15.8188 1.77315L15.8209 1.77568L15.8266 1.7828L15.8445 1.80519C15.8592 1.82368 15.8791 1.84928 15.9036 1.88166C15.9526 1.94636 16.0199 2.03834 16.0995 2.15482C16.2583 2.38724 16.4676 2.72011 16.6767 3.13086C17.0908 3.94443 17.526 5.10672 17.526 6.42082C17.526 7.74372 17.0537 8.96013 16.6124 9.81549C16.3887 10.2492 16.165 10.6052 15.9961 10.8545C15.9115 10.9794 15.84 11.0784 15.7884 11.1477C15.7626 11.1824 15.7417 11.2098 15.7265 11.2293L15.7082 11.2528L15.7024 11.2601L15.7004 11.2626L15.6996 11.2635C15.6996 11.2635 15.699 11.2643 15.0416 10.7648L15.699 11.2643C15.4153 11.6174 14.891 11.6799 14.5279 11.404C14.1652 11.1285 14.1006 10.6194 14.3832 10.2664L14.3839 10.2656L14.3845 10.2648L14.3937 10.253C14.4027 10.2413 14.4173 10.2223 14.4366 10.1964C14.4752 10.1445 14.5327 10.0651 14.6027 9.96176C14.7432 9.75443 14.9321 9.45393 15.1209 9.08796C15.5047 8.34397 15.8574 7.38839 15.8574 6.42082C15.8574 5.44444 15.5298 4.53559 15.1812 3.85064C15.0089 3.51212 14.8368 3.23889 14.7095 3.05252C14.646 2.9596 14.5942 2.88901 14.5597 2.8435C14.5425 2.82076 14.5297 2.80435 14.522 2.7946L14.5147 2.78541L14.5139 2.78442C14.2272 2.43477 14.2854 1.92507 14.6445 1.64536ZM11.9089 3.36249C12.2619 3.07454 12.7882 3.11941 13.0843 3.4627L12.7647 3.72338L12.4451 3.98407C13.0843 3.4627 13.0843 3.4627 13.0843 3.4627L13.085 3.46348L13.0857 3.46436L13.0875 3.46644L13.0921 3.47181L13.1051 3.48738C13.1155 3.49987 13.1291 3.51659 13.1456 3.53733C13.1784 3.57876 13.2227 3.6365 13.2746 3.70888C13.3779 3.85309 13.5135 4.05854 13.6491 4.31171C13.9152 4.80895 14.2072 5.53897 14.2072 6.37504C14.2072 7.23059 13.8702 7.99354 13.5744 8.50638C13.4225 8.76981 13.2709 8.98534 13.1562 9.13643C13.0986 9.21231 13.0496 9.27286 13.0136 9.31603C12.9956 9.33764 12.9807 9.35497 12.9696 9.36776L12.9558 9.38355L12.9511 9.38886L12.9493 9.39085L12.9485 9.39168C12.9485 9.39168 12.5362 9.04122 11.9089 8.50638L12.9479 9.3924C12.6441 9.72931 12.117 9.76297 11.7705 9.46759C11.4247 9.17278 11.3895 8.66165 11.6914 8.32471L11.6921 8.32386L11.6951 8.3205C11.6994 8.31555 11.7071 8.30652 11.7179 8.29365C11.7394 8.26786 11.7726 8.22694 11.8137 8.17285C11.8963 8.064 12.0078 7.90553 12.119 7.71276C12.3494 7.31319 12.5386 6.83487 12.5386 6.37504C12.5386 5.89574 12.3666 5.43027 12.1689 5.06087C12.0725 4.88071 11.9761 4.73506 11.9055 4.63648C11.8869 4.61052 11.8702 4.58801 11.856 4.56927L11.9089 3.36249ZM11.9089 3.36249L11.856 4.56927C11.8434 4.55265 11.8327 4.53897 11.8243 4.52838C11.8154 4.51714 11.8091 4.50946 11.8059 4.50556L11.8035 4.50278C11.5099 4.15951 11.5568 3.64969 11.9089 3.36249Z"
                fill="#7C7B7B"
              />
            </svg>
          </div>
        </div>
        <div className={style.exam__actions}>
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
              d="M20 1C20 0.447715 19.5523 0 19 0C18.4477 0 18 0.447715 18 1V19C18 19.5523 18.4477 20 19 20C19.5523 20 20 19.5523 20 19V1ZM8 5C8 4.44772 7.55228 4 7 4C6.44772 4 6 4.44772 6 5L6 19C6 19.5523 6.44772 20 7 20C7.55229 20 8 19.5523 8 19L8 5ZM2 11C2 10.4477 1.55228 10 1 10C0.447715 10 0 10.4477 0 11V19C0 19.5523 0.447715 20 1 20C1.55228 20 2 19.5523 2 19V11ZM13 10C13.5523 10 14 10.4477 14 11V19C14 19.5523 13.5523 20 13 20C12.4477 20 12 19.5523 12 19V11C12 10.4477 12.4477 10 13 10Z"
              fill="#52C3FF"
            />
          </svg>
          <div className={style.icon}>
            <svg
              width="20"
              height="20"
              viewBox="0 0 20 20"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              s
              <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M2 10C2 5.58172 5.58172 2 10 2C14.4183 2 18 5.58172 18 10C18 14.4183 14.4183 18 10 18C5.58172 18 2 14.4183 2 10ZM10 0C4.47715 0 0 4.47715 0 10C0 15.5228 4.47715 20 10 20C15.5228 20 20 15.5228 20 10C20 4.47715 15.5228 0 10 0ZM6.95818 5.2379C7.56342 4.56999 8.56884 3.81878 9.99999 3.81879C11.4239 3.8188 12.5389 4.54562 13.1841 5.6041C13.8143 6.63798 13.9861 7.95778 13.6764 9.22583C13.4048 10.3378 12.4984 11.2279 11.8469 11.7652C11.4965 12.0542 11.162 12.2871 10.9158 12.4477C10.792 12.5285 10.6888 12.592 10.6149 12.6363C10.5779 12.6585 10.5482 12.6759 10.5267 12.6883L10.5007 12.7032L10.4926 12.7078L10.4899 12.7094L10.4888 12.71L10.4883 12.7102C10.4881 12.7103 10.4879 12.7104 9.99999 11.8376L10.4879 12.7104C10.0059 12.9799 9.3966 12.8076 9.12711 12.3255C8.85786 11.8439 9.02967 11.2352 9.51079 10.9654L9.51097 10.9653L9.5114 10.9651L9.51165 10.9649L9.5117 10.9649L9.51204 10.9647L9.51262 10.9644L9.52575 10.9568C9.53853 10.9494 9.55911 10.9374 9.58648 10.921C9.6413 10.8881 9.72288 10.838 9.82316 10.7726C10.025 10.6409 10.2959 10.4519 10.5744 10.2222C11.1809 9.72202 11.6269 9.18762 11.7335 8.75135C11.9312 7.94182 11.8004 7.17673 11.4764 6.6451C11.1673 6.13808 10.6761 5.81879 9.99998 5.81879C9.33112 5.81878 8.82325 6.15817 8.4402 6.58088C8.25045 6.79028 8.10828 7.00327 8.01405 7.16429C7.96751 7.2438 7.93425 7.30792 7.91404 7.34886C7.90398 7.36926 7.89727 7.38366 7.89388 7.39107L7.89152 7.39631C7.67376 7.90096 7.08896 8.13557 6.58242 7.9204C6.0741 7.70447 5.83707 7.11734 6.053 6.60902L6.9734 7C6.053 6.60902 6.05316 6.60865 6.05333 6.60826L6.05367 6.60745L6.05443 6.60567L6.05623 6.60151L6.06092 6.59076L6.07476 6.55994C6.08597 6.5354 6.10122 6.50293 6.12059 6.46369C6.15926 6.38534 6.21479 6.27905 6.28793 6.15408C6.43306 5.90611 6.65379 5.5738 6.95818 5.2379ZM11 15.501C11 16.0529 10.5523 16.5003 9.99999 16.5003C9.44771 16.5003 8.99999 16.0529 8.99999 15.501C8.99999 14.9491 9.44771 14.5016 9.99999 14.5016C10.5523 14.5016 11 14.9491 11 15.501Z"
                fill="white"
              />
            </svg>
          </div>
          <div onClick={() => setISplaying(!isPlaying)} className={style.icon}>
            <svg
              width="18"
              height="18"
              viewBox="0 0 18 18"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M2.025 15.975V2.025H5.175V15.975H2.025ZM0 1.35C0 0.604414 0.604416 0 1.35 0H5.85C6.59558 0 7.2 0.604416 7.2 1.35V16.65C7.2 17.3956 6.59558 18 5.85 18H1.35C0.604417 18 0 17.3956 0 16.65V1.35ZM12.825 15.975V2.025H15.975V15.975H12.825ZM10.8 1.35C10.8 0.604414 11.4044 0 12.15 0H16.65C17.3956 0 18 0.604416 18 1.35V16.65C18 17.3956 17.3956 18 16.65 18H12.15C11.4044 18 10.8 17.3956 10.8 16.65V1.35Z"
                fill="white"
              />
              <path d="M2.025 2.025V15.975H5.175V2.025H2.025Z" fill="white" />
              <path
                d="M12.825 2.025V15.975H15.975V2.025H12.825Z"
                fill="white"
              />
            </svg>
          </div>
          <div className={style.icon}>
            <svg
              width="18"
              height="18"
              viewBox="0 0 18 18"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M4.71501 0C5.42186 0 5.99488 0.557009 5.99488 1.24411C5.99488 1.93122 5.42186 2.48823 4.71501 2.48823H2.98638C2.75076 2.48823 2.55975 2.6739 2.55975 2.90293V15.0971C2.55975 15.3261 2.75076 15.5118 2.98638 15.5118H4.71501C5.42186 15.5118 5.99488 16.0688 5.99488 16.7559C5.99488 17.443 5.42186 18 4.71501 18H2.98638C1.33705 18 0 16.7003 0 15.0971V2.90293C0 1.29969 1.33705 0 2.98638 0H4.71501ZM11.9713 12.5025L14.3371 10.2029H5.75043C5.04358 10.2029 4.47056 9.64584 4.47056 8.95874C4.47056 8.27164 5.04358 7.71463 5.75043 7.71463L14.3371 7.71463L11.9713 5.41494C11.4715 4.92909 11.4715 4.14136 11.9713 3.6555C12.4711 3.16964 13.2815 3.16965 13.7813 3.6555L17.1253 6.90606C18.2916 8.03972 18.2916 9.87776 17.1253 11.0114L13.7813 14.262C13.2815 14.7478 12.4711 14.7478 11.9713 14.262C11.4715 13.7761 11.4715 12.9884 11.9713 12.5025Z"
                fill="white"
              />
            </svg>
          </div>
        </div>
      </div>
    </div>
  );
};

export default TestHeader;
