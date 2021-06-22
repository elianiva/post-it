import {
  Box,
  Button,
  Center,
  FormControl,
  FormLabel,
  Heading,
  Input,
  Link,
  Text,
} from "@chakra-ui/react";
import React from "react";

export default function Login(): JSX.Element {
  return (
    <Box py="10rem">
      <Center>
        <Box as="form" minW="lg" p="6" rounded="md" bgColor="white" shadow="md">
          <Heading color="gray.800" textAlign="center" mb="2">
            Login
          </Heading>
          <Text color="gray.600" align="center" mb="8">
            You need to login first before posting stuff
          </Text>
          <FormControl isRequired mb="4">
            <FormLabel>Email</FormLabel>
            <Input type="email" placeholder="foobar@live.me" />
          </FormControl>
          <FormControl isRequired mb="8">
            <FormLabel>Password</FormLabel>
            <Input type="password" placeholder="********" />
          </FormControl>
          <Button colorScheme="blue" px="8" mx="auto" display="block" mb="4">
            Login
          </Button>
          <Text align="center" color="gray.600">
            Don&apos;t have an account?
            <Link color="blue.600"> Sign up here</Link>
          </Text>
        </Box>
      </Center>
    </Box>
  );
}
