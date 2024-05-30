import Scaffold from "@components/Scaffold";
import { Alert, Box, Typography } from "@mui/material";
import useAccountStore from "@store/account";
import { AccountInput, Credentials, IdentityInput } from "@types";
import { useCallback, useMemo, useState } from "react";
import SendEmailForm from "./SendEmailForm";
import PageCenter from "@components/PageCenter";
import { getAxiosError } from "@util/errors";
import { useNavigate, useSearchParams } from "react-router-dom";
import LoginForm from "./LoginForm";
import CreateIdentityForm from "./CreateIdentityForm";

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

interface Props {
  variant?: "login" | "signup" | "initiate";
}

const SignupPage: React.FC<Props> = ({ variant }) => {
  const navigate = useNavigate();
  const accountStore = useAccountStore();
  const [searchParams] = useSearchParams();

  const [error, setError] = useState("");
  const [emailSent, setEmailSent] = useState(false);

  const handleChange = useCallback(() => setError(""), []);

  const handleInitiate = useCallback((input: AccountInput) => {
    accountStore.initiateNewAccount(input).then(() => {
      setEmailSent(true);
    }).catch((err) => setError(getAxiosError(err, errMap)));
  }, [accountStore]);

  const handleLogin = useCallback((input: Credentials) => {
    accountStore.login(input).then((idn) => {
      navigate(`/account/${idn.accountId}`);
    }).catch((err) => setError(getAxiosError(err, { 404: "Account not found" })));
  }, [accountStore, navigate]);

  const handleCreateAccount = useCallback((input: IdentityInput) => {
    accountStore.createIdentity(input).then((idn) => {
      navigate(`/account/${idn.accountId}`);
    }).catch((err) => setError(getAxiosError(err, {})));
  }, [accountStore, navigate]);

  const title = useMemo(() => {
    switch (variant) {
      case "login":
        return "Log In";

      case "signup":
        return "Activate Account";

      case "initiate":
      default:
        return "Sign Up";
    }
  }, [variant]);

  const code = searchParams.get("code");
  const email = searchParams.get("email");

  const renderSwitch = () => {
    switch (variant) {
      case "login":
        return (
          <LoginForm onSubmit={handleLogin} onChange={handleChange} />
        );

      case "signup":
        if (!code || !email) {
          navigate("/signup");
          return null;
        }

        return (
          <CreateIdentityForm
            onSubmit={handleCreateAccount}
            code={code}
            email={email}
            onChange={handleChange} />
        );

      case "initiate":
      default:
        return emailSent ? (
          <>
            <Typography variant="h4" sx={styles.formHeader}>Email Sent</Typography>
            <Typography>Follow the link in your email to continue</Typography>
          </>
        ) : (
          <SendEmailForm onContinue={handleInitiate} onChange={handleChange} />
        );
    }
  };

  return (
    <Scaffold>
      <PageCenter>
        <Typography variant="h4" sx={styles.formHeader}>{title}</Typography>
        <Box sx={{ width: 400 }}>
          {renderSwitch()}
          {error && <Alert severity="error" sx={{ marginTop: 2, width: '100%' }}>{error}</Alert>}
        </Box>
      </PageCenter>
    </Scaffold>
  );
};

export default SignupPage;