import Image from "next/image";
import React from "react";

import style from "styles/home.module.scss";
import VideoPlayer from "../ui/VideoPlayer";

const HomeWorks = () => {
    const [show, setShow] = React.useState(false);
    return (
        <div className={style.home__description}>
            <div className={style.details}>
                <h1>How it works ?</h1>
                <p>
                    Li Europan lingues es membres del sam familie. Lor separat existentie
                    es un myth. Por scientie, musica, sport etc., li tot Europa usa li sam
                    vocabularium. Li lingues differe solmen in li grammatica, li
                    pronunciation e li plu commun vocabules. Omnicos directe al
                    desirabilit? de un nov lingua franca: on refusa continuar payar
                    custosi traductores. It solmen va esser necessi far uniform
                    grammatica, pronunciation e plu sommun paroles.
                </p>
            </div>
            <div className={style.slides}>
                <div className={style.calender}>
                    <Image src="/img/calandar.png" alt="" width={140} height={140}/>
                </div>
                <div className={style.play}>
                    {
                        show ? (
                            <VideoPlayer link={"https://www.youtube.com/watch?v=ysz5S6PUM-U"}/>

                        ) : (
                            // eslint-disable-next-line @next/next/no-img-element
                            <img src="/img/icons/play.svg" alt="play-icon" onClick={() => setShow(true)}/>

                        )
                    }
                </div>
                <div className={style.file}>
                    <Image src="/img/file.png" alt="file" width={156} height={156}/>
                </div>
            </div>
            <div className={style.mobile__details}>
                <p>
                    Li Europan lingues es membres del sam familie. Lor separat existentie
                    es un myth. Por scientie, musica, sport etc., li tot Europa usa li sam
                    vocabularium. Li lingues differe solmen in li grammatica, li
                    pronunciation e li plu commun vocabules. Omnicos directe al
                    desirabilit? de un nov lingua franca: on refusa continuar payar
                    custosi traductores. It solmen va esser necessi far uniform
                    grammatica, pronunciation e plu sommun paroles.
                </p>
            </div>
        </div>
    );
};

export default HomeWorks;
