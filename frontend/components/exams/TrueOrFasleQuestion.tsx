import React from "react";
import {
  FormControl,
  FormControlLabel,
  FormLabel,
  Radio,
  RadioGroup,
} from "@mui/material";

import style from "styles/exams/questions.module.scss";

interface Props {
  title: string;
  count: number;
}

const TrueOrFalseQuestion: React.FC<Props> = ({ title, count }) => {
  return (
    <FormControl
      sx={{
        flexDirection: "row",
        alignItems: "center",
        minWidth: "100%",
        justifyContent: "space-between",
      }}
      className={style.trueOrFalse}
    >
      <FormLabel
        sx={{ width: "100%", fontSize: 14 }}
        className={style.label}
        id="demo-radio-buttons-group-label"
      >
        {count + " - " + title}
      </FormLabel>
      <RadioGroup
        sx={{ width: "250px", flexDirection: "row", padding: "0 1em" }}
        className={style.actions}
        aria-labelledby="demo-radio-buttons-group-label"
        name="radio-buttons-group"
      >
        <FormControlLabel value="true" control={<Radio />} label="True" />
        <FormControlLabel value="false" control={<Radio />} label="False" />
      </RadioGroup>
    </FormControl>
  );
};

export default TrueOrFalseQuestion;
