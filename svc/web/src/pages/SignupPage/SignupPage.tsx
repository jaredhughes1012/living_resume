import Scaffold from "@components/Scaffold";
import { Alert, Typography } from "@mui/material";
import useAccountStore from "@store/account";
import { AccountInput } from "@types";
import { useCallback, useState } from "react";
import SendEmailForm from "./SendEmailForm";
import PageCenter from "@components/PageCenter";
import { getAxiosError } from "@util/errors";

const styles = {
  pageFrame: {
    width: "100vw",
    height: "100vh",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "center",
  },
  formHeader: {
    marginBottom: 4,
  },
};

const errMap = {
  409: "An account with that email already exists",
};

interface Props { }

const SignupPage: React.FC<Props> = () => {
  const accountStore = useAccountStore();

  const [error, setError] = useState("");
  const [emailSent, setEmailSent] = useState(false);

  const handleChange = useCallback(() => setError(""), []);

  const handleContinue = useCallback((input: AccountInput) => {
    accountStore.initiateNewAccount(input).then(() => {
      setEmailSent(true);
    }).catch((err) => setError(getAxiosError(err, errMap)));
  }, [accountStore]);

  return (
    <Scaffold>
      <PageCenter>
        {emailSent ? (
          <>
            <Typography variant="h4" sx={styles.formHeader}>Email Sent</Typography>
            <Typography>Follow the link in your email to continue</Typography>
          </>
        ) : (
          <>
            <Typography variant="h4" sx={styles.formHeader}>Create an account</Typography>
            <SendEmailForm
              onContinue={handleContinue}
              onChange={handleChange} />
            {error && <Alert severity="error" sx={{ marginTop: 2 }}>{error}</Alert>}
          </>
        )}
      </PageCenter>
    </Scaffold>
  );
};

export default SignupPage;