import { Box } from "@mui/material";

const styles = {
  pageFrame: {
    flexGrow: 1,
    height: "100vh",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "center",
  },
}

interface Props {
  children?: React.ReactNode;
  vertical?: boolean;
}

// Fills the entire page and centers its children horizontally. Optional to also center vertically
const PageCenter: React.FC<Props> = ({ children }) => {
  return (
    <Box sx={styles.pageFrame}>
      {children}
    </Box>
  );
};

export default PageCenter;