import { Box, Button, Theme, Typography } from "@mui/material";
import { useCallback } from "react";
import { useNavigate } from "react-router-dom";

const styles = {
  container: (theme: Theme) => ({
    padding: 2,
    display: "flex",
    flexDirection: "column",
    height: "100%",
    minHeight: "100vh",
    width: 250,
    backgroundColor: theme.palette.background.paper,
  }),

  navButton: {
    marginTop: 1,
  }
}

interface Props { }

const NavRail: React.FC<Props> = () => {
  const navigate = useNavigate();

  const handleSignUp = useCallback(() => {
    navigate("/signup");
  }, [navigate]);

  return (
    <Box sx={styles.container}>
      <Box sx={{ flexGrow: 1 }} />
      <Typography>Log In or Sign Up</Typography>
      <Typography variant="caption">Create an account to set up your own living resume</Typography>
      <Button
        variant="contained"
        color="primary"
        sx={styles.navButton}
        onClick={handleSignUp}>
        Sign Up
      </Button>

      <Button
        variant="contained"
        color="inherit"
        sx={styles.navButton}>
        Log In
      </Button>
    </Box>
  );
};

export default NavRail;