import React from "react";
import { Recorder } from "react-voice-recorder";
import "react-voice-recorder/dist/index.css";

class AudioRecorder extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      url: null,
      blob: null,
      chunks: null,
      duration: {
        h: 0,
        m: 0,
        s: 0,
      },
    };
  }

  handleAudioStop(data) {
    console.log(data);
    this.setState(data);
  }

  handleReset() {
    const reset = {
      url: null,
      blob: null,
      chunks: null,
      duration: {
        h: 0,
        m: 0,
        s: 0,
      },
    };
    this.setState(reset);
  }

  handleAudioUpload(file) {
    console.log(file);
  }

  render() {
    return (
      <Recorder
        record={true}
        audioURL={this.state.url}
        showUIAudio
        handleAudioStop={(data) => this.handleAudioStop(data)}
        handleOnChange={(value) => this.handleOnChange(value, "firstname")}
        handleAudioUpload={(data) => this.handleAudioUpload(data)}
        handleRest={() => this.handleReset()}
      />
    );
  }
}

export default AudioRecorder;
