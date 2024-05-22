import { IconButton, TextField } from "@mui/material";
import { ChatInput } from "@types";
import { useCallback, useState } from "react";
import IconArrowUp from "@mui/icons-material/ArrowUpward";

const styles = {
  inputBar: {
    width: 800,
  }
}

interface Props {
  onSubmit: (input: ChatInput) => void;
}

// Input used to communicate with the chat api
const InputBar: React.FC<Props> = ({ onSubmit }) => {
  const [value, setValue] = useState("");

  const handleChange = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
    setValue(event.target.value);
  }, []);

  const handleSubmit = useCallback(() => {
    if (!value) return;

    onSubmit({ message: value });
    setValue("");
  }, [onSubmit, value]);

  const handleKeyDown = useCallback((event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === "Enter") {
      event.stopPropagation();
      handleSubmit();
    }
  }, [handleSubmit]);

  return (
    <TextField
      label="Ask a question"
      sx={styles.inputBar}
      value={value}
      onChange={handleChange}
      multiline
      maxRows={5}
      onKeyDown={handleKeyDown}
      InputProps={{
        endAdornment: (
          <IconButton size="small" sx={{ margin: -1 }} onClick={handleSubmit}>
            <IconArrowUp />
          </IconButton>
        )
      }} />
  )
};

export default InputBar;