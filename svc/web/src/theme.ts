import { Shadows, createTheme } from "@mui/material";

const theme = createTheme({
  palette: {
    mode: "dark",
    primary: {
      main: "#4caf50",
    },
    secondary: {
      main: "#00e676",
    },
    background: {
      paper: "#121212",
      default: "#232323",
    }
  },

  typography: {
    button: {
      textTransform: "none",
    }
  },

  shape: {
    borderRadius: 12,
  },

  shadows: Array<string>(25).fill("none") as Shadows,
});

export default theme;