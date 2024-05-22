import { Box, Theme } from "@mui/material";

const styles = {
  line: {
    flexGrow: 1,
    borderBottom: (theme: Theme) => `1px solid ${theme.palette.divider}`,
  },
  divider: {
    display: "flex",
    alignItems: "center",
    flexDirection: "row",
    width: "100%",
    margin: "10px 0",
  },
  orText: {
    margin: "0 10px",
  },
}

interface Props { }

const OrDivider: React.FC<Props> = () => {
  return (
    <Box sx={styles.divider}>
      <Box sx={styles.line} />
      <Box sx={styles.orText}>OR</Box>
      <Box sx={styles.line} />
    </Box>
  );
};

export default OrDivider;