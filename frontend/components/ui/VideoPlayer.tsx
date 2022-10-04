import React from "react"
import ReactPlayer from 'react-player/youtube'

interface Props {
    link: string
}

const VideoPlayer: React.FC<Props> = ({link}) => {
    return (
        <ReactPlayer url={link} width={"100%"} height={"100%"} style={{borderRadius: "30px!important"}}
                     controls={true}/>
    );
};

export default VideoPlayer
