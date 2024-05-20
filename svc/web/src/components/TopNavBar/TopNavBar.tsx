import { Box, Link, Popover, PopoverContent, PopoverTrigger, Stack, useColorModeValue } from "@chakra-ui/react";
import NavBarBase from "./NavBarBase";

interface NavItem {
  link: string;
  label: string;
  children?: NavItem[];
}

const navItems: NavItem[] = [
  {
    link: "/",
    label: "Home",
    children: [
      {
        link: "/about",
        label: "About"
      },
    ]
  }
];

interface NavItemProps {
  item: NavItem;
  bgColor: string;
  linkHoverColor: string;
}

const NavItem: React.FC<NavItemProps> = ({ item, bgColor, linkHoverColor }) => {
  return (
    <Box key={item.link}>
      <Popover trigger="hover" placement="bottom-start">
        <PopoverTrigger>
          <Link
            href={item.link}
            _hover={{
              textDecoration: 'none',
              color: linkHoverColor,
            }}>
            {item.label}
          </Link>
        </PopoverTrigger>
        {item.children && (
          <PopoverContent
            border={0}
            boxShadow={'xl'}
            bg={bgColor}
            p={4}
            rounded={'xl'}
            minW={'sm'}>
            <Stack>
              {item.children.map((child) => (
                <NavItem key={child.label} item={child} bgColor={bgColor} linkHoverColor={linkHoverColor} />
              ))}
            </Stack>
          </PopoverContent>
        )}
      </Popover>
    </Box>
  );
}

interface Props { }

const TopNavBar: React.FC<Props> = () => {
  const popoverBgColor = useColorModeValue('white', 'gray.800');
  const linkHoverColor = useColorModeValue('gray.800', 'white');

  return (
    <NavBarBase>
      <Stack direction="row" spacing={4}>
        {navItems.map((item) => (
          <NavItem
            key={item.label}
            item={item}
            bgColor={popoverBgColor}
            linkHoverColor={linkHoverColor} />
        ))}
      </Stack>
    </NavBarBase>
  );
};

export default TopNavBar;