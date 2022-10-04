import CustomInput from "components/ui/CustomInput";
import Image from "next/image";
import React from "react";

import style from "styles/dashboard/profile.module.scss";
import {RootStateOrAny, useSelector} from "react-redux";

const Inputs = () => {
    const {email , userName} = useSelector((state: RootStateOrAny) => state.auth)
    const [_name, setName] = React.useState<string | number>("");
    const [_email, setEmail] = React.useState<string | number>("");
    const [password, setPassword] = React.useState<string | number>("");
    const [repeatPassword, setRepeatPassword] = React.useState<string | number>(
        ""
    );
    return (
        <div className={style.profile__inputs}>
            <CustomInput
                Icon={() => (
                    <Image
                        src="/img/icons/person.svg"
                        alt="image"
                        width={10}
                        height={14}
                    />
                )}
                setValue={setName}
                value={userName}
                title="Full Name"
            />
            <CustomInput
                Icon={() => (
                    <Image
                        src="/img/icons/email.svg"
                        alt="image"
                        width={10}
                        height={14}
                    />
                )}
                setValue={setEmail}
                value={email}
                title="Email"
            />
            <CustomInput
                Icon={() => (
                    <Image src="/img/icons/lock.svg" alt="lock" width={10} height={14}/>
                )}
                setValue={setPassword}
                value={password}
                title="Password"
            />
            <CustomInput
                Icon={() => (
                    <Image src="/img/icons/lock.svg" alt="image" width={10} height={14}/>
                )}
                setValue={setRepeatPassword}
                value={repeatPassword}
                title="Repeat Password"
            />
        </div>
    );
};

export default Inputs;
