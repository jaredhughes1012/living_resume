import OrDivider from "@components/OrDivider";
import Scaffold from "@components/Scaffold";
import { Box, Button, Link, TextField, Typography } from "@mui/material";
import { Link as RRLink } from "react-router-dom";

const styles = {
  pageFrame: {
    width: "100vw",
    height: "100vh",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "center",
  },
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

interface Props { }

const SignupPage: React.FC<Props> = () => {
  return (
    <Scaffold>
      <Box sx={styles.pageFrame}>
        <Box sx={styles.form}>
          <Typography variant="h4" sx={styles.formHeader}>Create an account</Typography>
          <TextField
            sx={styles.formItem}
            label="Email"
            variant="outlined" />
          <Button variant="contained" color="primary" sx={styles.formItem}>Continue</Button>
          <Typography sx={styles.loginText}>Already have an account? <Link component={RRLink} to="/login">Log In</Link></Typography>
          <OrDivider />

          <Button variant="contained" color="inherit" sx={styles.formItem}>Other Signup Buttons Here</Button>
        </Box>
      </Box>
    </Scaffold>
  );
};

export default SignupPage;