import { Box, Button, TextField } from "@mui/material";
import { IdentityInput } from "@types";
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
  onSubmit: (input: IdentityInput) => void;
  code: string;
  email: string;
}

const CreateIdentityForm: React.FC<Props> = ({ onSubmit, code, email }) => {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [accountId, setAccountId] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = useCallback(() => {
    const input: IdentityInput = {
      activationCode: code,
      firstName,
      lastName,
      accountId,
      credentials: {
        email,
        password,
      },
    };

    onSubmit(input);
  }, [code, email, firstName, lastName, password, accountId, onSubmit]);

  const handleFirstNameChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => setFirstName(e.target.value), []);
  const handleLastNameChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => setLastName(e.target.value), []);
  const handleAccountIdChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => setAccountId(e.target.value), []);
  const handlePasswordChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => setPassword(e.target.value), []);
  const handleEnter = useCallback((e: React.KeyboardEvent) => e.key === "Enter" && handleSubmit(), [handleSubmit]);

  return (
    <Box sx={{ width: 400 }}>
      <Box sx={[styles.nameRow, styles.row]}>
        <TextField
          label="First Name"
          value={firstName}
          onChange={handleFirstNameChange}
          sx={{ marginRight: 1 }}
          onKeyDown={handleEnter} />

        <TextField
          label="Last Name"
          value={lastName}
          onChange={handleLastNameChange}
          sx={{ marginLeft: 1 }}
          onKeyDown={handleEnter} />
      </Box>

      <TextField
        label="Account ID"
        value={accountId}
        onChange={handleAccountIdChange}
        helperText="This ID will be visible in the URL you share with others"
        sx={styles.row}
        type="text"
        onKeyDown={handleEnter} />

      <TextField
        label="Email"
        value={email}
        disabled
        sx={styles.row} />

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
        Create Account
      </Button>
    </Box>
  );
};

export default CreateIdentityForm;