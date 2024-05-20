import { AppBar, Box, Toolbar, Typography } from "@mui/material";

// const navItems: NavItem[] = [
//   {
//     link: "/",
//     label: "Home",
//   }
// ];

interface Props { }

const TopNavBar: React.FC<Props> = () => {
  return (
    <AppBar>
      <Toolbar>
        <Typography variant="h6" component="div">
          Jared Hughes
        </Typography>
        <Box sx={{ flexGrow: 1 }} />
      </Toolbar>
    </AppBar>
  );
};

export default TopNavBar;