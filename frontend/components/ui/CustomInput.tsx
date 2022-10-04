import React, {FC} from "react";
import style from "styles/auth.module.scss";

interface Props {
    setValue: (event: string) => void;
    value: string | number;
    title: string;
    Icon: React.FC;
    inputType?: string;
    register?: any;
}

const AuthInput: FC<Props> = ({setValue, value, title, Icon, inputType}) => {
    return (
        <div style={{width: "100%"}} className={style.input__container}>
      <span className={style.icon}>
        <Icon/>
      </span>
            <input
                readOnly={true}
                style={{width: "100%"}}
                className={style.input}
                type={inputType ? inputType : "text"}
                placeholder={title}
                value={value}
                onChange={(event) => setValue(event.target.value)}
            />
        </div>
    );
};

export default AuthInput;
