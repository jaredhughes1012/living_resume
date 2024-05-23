import PageCenter from "@components/PageCenter";
import Scaffold from "@components/Scaffold";
import CreateIdentityForm from "./CreateIdentityForm";
import useAccountStore from "@store/account";
import { useCallback } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { IdentityInput } from "@types";
import { Typography } from "@mui/material";

interface Props { }

const CreateAccountPage: React.FC<Props> = () => {
  const accountStore = useAccountStore();
  const navigate = useNavigate();
  const [params] = useSearchParams();

  const handleSubmit = useCallback((input: IdentityInput) => {
    accountStore.createIdentity(input).then((idn) => {
      navigate(`/account/${idn.accountId}`);
    });
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
          code={code}
          email={email} />
      </PageCenter>
    </Scaffold>
  );
};

export default CreateAccountPage;