import PageCenter from "@components/PageCenter";
import Scaffold from "@components/Scaffold";
import CreateIdentityForm from "./CreateIdentityForm";
import useAccountStore from "@store/account";
import { useCallback, useState } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { IdentityInput } from "@types";
import { Alert, Typography } from "@mui/material";
import { getAxiosError } from "@util/errors";

interface Props { }

const CreateAccountPage: React.FC<Props> = () => {
  const accountStore = useAccountStore();
  const navigate = useNavigate();
  const [params] = useSearchParams();

  const [error, setError] = useState("");

  const handleChange = useCallback(() => setError(""), []);

  const handleSubmit = useCallback((input: IdentityInput) => {
    accountStore.createIdentity(input).then((idn) => {
      navigate(`/account/${idn.accountId}`);
    }).catch((err) => setError(getAxiosError(err, {})));
  }, [accountStore, navigate]);

  const code = params.get("code");
  const email = params.get("email");

  if (!code || !email) {
    navigate("/signup");
    return null;
  }

  return (
    <Scaffold>
      <PageCenter>
        <Typography variant="h4" sx={{ marginBottom: 2 }}>Finish creating your account</Typography>
        <CreateIdentityForm
          onSubmit={handleSubmit}
          onChange={handleChange}
          code={code}
          email={email} />

        {error && <Alert severity="error" sx={{ marginTop: 2 }}>{error}</Alert>}
      </PageCenter>
    </Scaffold>
  );
};

export default CreateAccountPage;