import OrDivider from "@components/OrDivider";
import { useChangePropagator, useEnterHandler } from "@util/hooks";
import { Box, Button, Link, TextField, Typography } from "@mui/material";
import { AccountInput } from "@types";
import { useCallback, useState } from "react";
import { Link as RRLink } from "react-router-dom";

const styles = {
  form: {
    width: 350,
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  },
  formItem: {
    width: "100%",
    marginTop: 2,
  },
  formHeader: {
    marginBottom: 4,
  },
  loginText: {
    marginTop: 2,
  }
};

interface Props {
  onContinue: (input: AccountInput) => void;
  onChange: () => void;
}

// Form used to send an activation email to the user
const SendEmailForm: React.FC<Props> = ({ onContinue, onChange }) => {
  const [email, setEmail] = useState('');

  const handleContinue = useCallback(() => onContinue({ email }), [onContinue, email]);
  const handleEmailChange = useChangePropagator(setEmail, onChange);
  const handleEnter = useEnterHandler(handleContinue);

  return (
    <Box sx={styles.form}>
      <TextField
        sx={styles.formItem}
        value={email}
        onChange={handleEmailChange}
        onKeyDown={handleEnter}
        label="Email"
        variant="outlined" />
      <Button
        variant="contained"
        color="primary"
        onClick={handleContinue}
        sx={styles.formItem}>
        Continue
      </Button>
      <Typography sx={styles.loginText}>Already have an account? <Link component={RRLink} to="/login">Log In</Link></Typography>
      <OrDivider />

      <Button variant="contained" color="inherit" sx={styles.formItem}>
        Other Signup Buttons Here
      </Button>
    </Box>
  );
}

export default SendEmailForm;