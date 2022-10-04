import { login } from "app/ActionCreators";
import AuthInput from "components/ui/CustomInput";
import HomeLayout from "components/ui/HomeLayout";
import Navbar from "components/ui/Navbar";
import Image from "next/image";
import style from "styles/auth.module.scss";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import React, { useEffect } from "react";
import { useForm } from "react-hook-form";
import { RootStateOrAny, useDispatch, useSelector } from "react-redux";
import Footer from "../components/ui/Footer";

const SignIn = () => {
  const regex =
    /^(([^<>()[\]\.,;:\s@\"]+(\.[^<>()[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i;
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<any>({
    mode: "onBlur",
  });
  const dispatch = useDispatch();
  const onSubmit = () => {
    dispatch(login(watch().email, watch().password));
  };
  return (
    <HomeLayout>
      <Navbar
        links={[
          { link: "/signup", title: "Sign up" },
          { link: "/signin", title: "Sign in" },
        ]}
      />
      <ToastContainer theme="colored" />
      <div className={style.auth}>
        {/* image section */}
        <div className={style.icons}>
          <div className={style.calender}>
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
          <h1>Lets Sign in!</h1>
          <h4>
            We help you to find the best courses <br /> and virtual exam to
            learn anything!
          </h4>
          <form className={style.inputs}>
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
            {errors.password && (
              <p className={style.error__message}>
                {errors?.password?.message?.toString()}
              </p>
            )}
            <button onClick={handleSubmit(onSubmit)} className={style.button}>
              Sign in
            </button>
          </form>
          <div className={style.forgot}>
            <p>Forgot password ?</p>
            <button className={style.facebook}>Sign in with Facebook</button>
            <button className={style.google}>Sign in with Google</button>
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
