import { Box, Button, TextField } from "@mui/material";
import { IdentityInput } from "@types";
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
  onSubmit: (input: IdentityInput) => void;
  onChange: () => void;
  code: string;
  email: string;
}

const CreateIdentityForm: React.FC<Props> = ({ onSubmit, code, email, onChange }) => {
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

  const handleFirstNameChange = useChangePropagator(setFirstName, onChange);
  const handleLastNameChange = useChangePropagator(setLastName, onChange);
  const handleAccountIdChange = useChangePropagator(setAccountId, onChange);
  const handlePasswordChange = useChangePropagator(setPassword, onChange);

  const handleEnter = useEnterHandler(handleSubmit);

  return (
    <>
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
    </>
  );
};

export default CreateIdentityForm;