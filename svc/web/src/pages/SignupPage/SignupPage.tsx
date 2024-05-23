import Scaffold from "@components/Scaffold";
import { Typography } from "@mui/material";
import useAccountStore from "@store/account";
import { AccountInput } from "@types";
import { useCallback, useState } from "react";
import SendEmailForm from "./SendEmailForm";
import PageCenter from "@components/PageCenter";

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

interface Props { }

const SignupPage: React.FC<Props> = () => {
  const accountStore = useAccountStore();

  const [emailSent, setEmailSent] = useState(false);

  const handleContinue = useCallback((input: AccountInput) => {
    accountStore.initiateNewAccount(input).then(() => {
      setEmailSent(true);
    });
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
            <SendEmailForm onContinue={handleContinue} />
          </>
        )}
      </PageCenter>
    </Scaffold>
  );
};

export default SignupPage;