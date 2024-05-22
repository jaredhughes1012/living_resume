import NavRail from "@components/NavRail";
import { Box } from "@mui/material";
import InputBar from "./InputBar";
import Scaffold from "@components/Scaffold";
import { useCallback } from "react";
import { ChatInput } from "@types";

const styles = {
  chatFrame: {
    flexGrow: 1,
    height: "100vh",
    display: "flex",
    flexDirection: "column",
    padding: 4,
    alignItems: "center",
  }
}

interface Props { }

const LandingPage: React.FC<Props> = () => {
  const handleChatSubmit = useCallback((input: ChatInput) => console.log(input), []);

  return (
    <Scaffold>
      <NavRail />
      <Box sx={styles.chatFrame}>
        <Box sx={{ flexGrow: 1 }} />
        <InputBar onSubmit={handleChatSubmit} />
      </Box>
    </Scaffold>
  );
};

export default LandingPage;