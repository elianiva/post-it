import {
  Box,
  Button,
  Center,
  Container,
  Flex,
  Heading,
  Spacer,
  Link,
} from "@chakra-ui/react";
import React from "react";
import { Link as RouteLink, useLocation } from "react-router-dom";

export default function Navbar(): JSX.Element {
  const location = useLocation();
  const isLoginPage = location.pathname === "/login";

  return (
    <Box
      position="absolute"
      left="0"
      right="0"
      top="0"
      p="4"
      bgColor="white"
      borderBottom="2px"
      borderColor="gray.200"
      zIndex="10"
    >
      <Container maxW="container.lg">
        <Flex align="center">
          <Center>
            <Heading color="gray.700" size="lg">
              Post-It
            </Heading>
          </Center>
          <Spacer />
          <Flex gridColumnGap="1rem" align="center">
            <Box as="span" fontWeight="600" color="gray.600">
              <Link as={RouteLink} to="/">
                Home
              </Link>
            </Box>
            <Box as="span" fontWeight="600" color="gray.600">
              About
            </Box>
            <Button colorScheme="blue" outline="none">
              <Link to={isLoginPage ? "/signup" : "login"}>
                {isLoginPage ? "Sign Up" : "Login"}
              </Link>
            </Button>
          </Flex>
        </Flex>
      </Container>
    </Box>
  );
}
