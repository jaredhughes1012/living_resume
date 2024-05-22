import { Box } from "@mui/material";

const styles = {
  container: {
    display: "flex",
    minHeight: "100vh",
    minWidth: "100vw",
    flexDirection: "row",
  }
};
interface Props {
  children?: React.ReactNode;
}

// Primary layout for the page
const Scaffold: React.FC<Props> = ({ children }) => {
  return (
    <Box sx={styles.container}>
      {children}
    </Box>
  );
};

export default Scaffold;