import React, { FC, useEffect, useRef } from "react";

interface Props {
  isPlaying: boolean;
}

const AudioPlayer: FC<Props> = ({ isPlaying }) => {
  const playRef = useRef<any>();
  useEffect(() => {
    console.log(isPlaying);
    if (isPlaying) {
      playRef.current.play();
      return;
    }
    playRef.current.pause();
  }, [isPlaying]);
  return (
    <audio
      style={{ display: "none" }}
      autoPlay={isPlaying}
      controls
      ref={playRef}
    >
      <source
        src="https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3"
        type="audio/mpeg"
      />
    </audio>
  );
};

export default AudioPlayer;
