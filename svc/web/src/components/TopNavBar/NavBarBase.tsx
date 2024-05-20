import { Box, Flex, useColorModeValue } from "@chakra-ui/react";

interface Props {
  children?: React.ReactNode;
}

const NavBarBase: React.FC<Props> = ({ children }) => {
  return (
    <Box>
      <Flex
        bg={useColorModeValue('white', 'gray.800')}
        color={useColorModeValue('gray.600', 'white')}
        minH='60px'
        py={{ base: 2 }}
        px={{ base: 4 }}
        borderBottom={1}
        borderStyle='solid'
        align='center'
        borderColor={useColorModeValue('gray.200', 'gray.900')}>
        {children}
      </Flex>
    </Box>
  );
};

export default NavBarBase;