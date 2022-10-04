import HomeLayout from "components/ui/HomeLayout";
import Navbar from "components/ui/Navbar";
import Image from "next/image";
import React, { useEffect } from "react";
import { useForm } from "react-hook-form";
import PasswordStrengthBar from "react-password-strength-bar";
import style from "styles/auth.module.scss";
import "react-toastify/dist/ReactToastify.css";
import { toast, ToastContainer } from "react-toastify";
import { signUp } from "app/ActionCreators";
import { RootStateOrAny, useDispatch, useSelector } from "react-redux";
import Footer from "../components/ui/Footer";

const SignIn = () => {
  const regex =
    /^(([^<>()[\]\.,;:\s@\"]+(\.[^<>()[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s`@`\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i;
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<any>({
    mode: "onBlur",
  });

  const [passwordStrengthBar, setPasswordStrengthBar] = React.useState(0);
  const dispatch = useDispatch();
  const onSubmit = () => {
    dispatch(
      signUp(watch().name, watch().family, watch().email, watch().password)
    );
  };
  return (
    <HomeLayout>
      <Navbar
        links={[
          { link: "/signup", title: "Sign up" },
          { link: "/signin", title: "Sign in" },
        ]}
      />
      <ToastContainer theme={"colored"} />
      <div className={style.auth}>
        {/* image section */}
        <div className={style.icons}>
          <div className={style.calender}>
            1
            <Image
              src={"/img/calandar.png"}
              alt="calender"
              width={140}
              height={140}
            />
          </div>
          <div className={style.cap}>
            <Image src={"/img/cap.png"} alt="cap" width={295} height={250} />
          </div>
          <div className={style.file}>
            <Image src={"/img/file.png"} alt="file" width={140} height={140} />
          </div>
        </div>
        {/* action section */}
        <div className={style.details}>
          <h1>Lets Sign up!</h1>
          <h4>
            We help you to find the best courses <br /> and virtual exam to
            learn anything!
          </h4>
          <form className={style.inputs}>
            {/* ################################################ Name Input ################################################*/}
            <div className={style.input__container}>
              <span className={style.icon}>
                <Image
                  src="/img/icons/person.svg"
                  alt="user"
                  width={10}
                  height={14}
                />
              </span>

              <input
                className={style.input}
                type={"text"}
                placeholder={"Name"}
                {...register("name", {
                  required: "this field is required*",
                  minLength: {
                    value: 3,
                    message: "minimum length must be at least 3 characters",
                  },
                })}
              />
            </div>
            {errors.name && (
              <p className={style.error__message}>
                {errors?.name?.message?.toString()}
              </p>
            )}
            {/*################################################ Family Input ################################################*/}
            <div className={style.input__container}>
              <span className={style.icon}>
                <Image
                  src="/img/icons/person.svg"
                  alt="user"
                  width={10}
                  height={14}
                />
              </span>

              <input
                className={style.input}
                type={"text"}
                placeholder={"Family"}
                {...register("family", {
                  required: "this field is required*",
                  minLength: {
                    value: 3,
                    message: "minimum length must be at least 3 characters",
                  },
                })}
              />
            </div>
            {errors.name && (
              <p className={style.error__message}>
                {errors?.name?.message?.toString()}
              </p>
            )}
            {/* ################################################ Email Input ################################################ */}
            <div className={style.input__container}>
              <span className={style.icon}>
                <Image
                  src="/img/icons/email.svg"
                  alt="user"
                  width={10}
                  height={14}
                />
              </span>
              <input
                className={style.input}
                type={"text"}
                {...register("email", {
                  required: "this field is required*",
                  pattern: {
                    value: regex,
                    message: "Please enter a valid email address",
                  },
                  minLength: {
                    value: 6,
                    message: "minimum length must be at least 6 characters",
                  },
                })}
                placeholder={"Email"}
              />
            </div>
            {errors.email && (
              <p className={style.error__message}>
                {errors?.email?.message?.toString()}
              </p>
            )}
            {/* ################################################ Password Input ################################################ */}
            <div className={style.input__container}>
              <span className={style.icon}>
                <Image
                  src="/img/icons/lock.svg"
                  alt="user"
                  width={10}
                  height={14}
                />
              </span>
              <input
                className={style.input}
                type={"password"}
                {...register("password", {
                  required: "this field is required*",
                  minLength: {
                    value: 8,
                    message: "minimum length must be at least 8 characters",
                  },
                })}
                placeholder={"Password"}
              />
            </div>
            <PasswordStrengthBar
              password={watch("password")}
              onChangeScore={(value) => setPasswordStrengthBar(value)}
              minLength={8}
              scoreWords={["weak", "good", "perfect"]}
            />
            {errors.password && (
              <p className={style.error__message}>
                {errors?.password?.message?.toString()}
              </p>
            )}
            {/*################################################ Repeat Password Input ################################################ */}
            <div className={style.input__container}>
              <span className={style.icon}>
                <Image
                  src="/img/icons/lock.svg"
                  alt="user"
                  width={10}
                  height={14}
                />
              </span>
              <input
                className={style.input}
                type={"password"}
                {...register("repeatPassword", {
                  required: "this field is required*",
                  minLength: {
                    value: 8,
                    message: "minimum length must be at least 8   characters",
                  },

                  validate: (val: string) => {
                    if (watch("password") != val) {
                      return "your password isn't same";
                    } else if (passwordStrengthBar < 2) {
                      return "your password is too weak!";
                    }
                  },
                })}
                placeholder={"Repeat Password"}
              />
            </div>
            {errors.repeatPassword && (
              <p className={style.error__message}>
                {errors?.repeatPassword?.message?.toString()}
              </p>
            )}
            <p className={style.rule}>
              By Registering, You confirm that you accept our
              <a> Terms of Use </a>
              and
              <a> Privacy Policy</a>
            </p>
            <button onClick={handleSubmit(onSubmit)} className={style.button}>
              Sign up
            </button>
          </form>
          <div className={style.forgot}>
            <button className={style.facebook}>Sign up with Facebook</button>
            <button className={style.google}>Sign up with Google</button>
            <div className={style.forgot__details}>
              <span>Hot</span>
              <p>Over 1000 students are taking the IELTS course and Exams</p>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </HomeLayout>
  );
};

export default SignIn;
