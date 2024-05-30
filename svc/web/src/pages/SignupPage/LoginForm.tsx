import { Button, TextField } from "@mui/material";
import { Credentials } from "@types";
import { useChangePropagator, useEnterHandler } from "@util/hooks";
import { useCallback, useState } from "react";

const styles = {
  nameRow: {
    display: "flex",
    flexDirection: "row",
  },
  row: {
    marginTop: 2,
    width: "100%",
  }
};

interface Props {
  onSubmit: (input: Credentials) => void;
  onChange: () => void;
}

const LoginForm: React.FC<Props> = ({ onSubmit, onChange }) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = useCallback(() => {
    const input: Credentials = {
      email,
      password,
    };

    onSubmit(input);
  }, [email, password, onSubmit]);

  const handleEmailChange = useChangePropagator(setEmail, onChange);
  const handlePasswordChange = useChangePropagator(setPassword, onChange);
  const handleEnter = useEnterHandler(handleSubmit);

  return (
    <>
      <TextField
        label="Email"
        value={email}
        onChange={handleEmailChange}
        sx={styles.row}
        onKeyDown={handleEnter} />

      <TextField
        label="Password"
        value={password}
        onChange={handlePasswordChange}
        type="password"
        sx={styles.row}
        onKeyDown={handleEnter} />

      <Button
        variant="contained"
        color="primary"
        onClick={handleSubmit}
        size="large"
        sx={styles.row}>
        Log In
      </Button>
    </>
  );
};

export default LoginForm;