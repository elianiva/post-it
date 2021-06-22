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

export default function Signup(): JSX.Element {
  return (
    <Box py="8rem">
      <Center>
        <Box as="form" minW="lg" p="6" rounded="md" bgColor="white" shadow="md">
          <Heading color="gray.800" textAlign="center" mb="2">
            Sign Up
          </Heading>
          <Text color="gray.600" align="center" mb="8">
            Create an account to start posting stuff
          </Text>
          <FormControl isRequired mb="4">
            <FormLabel>Username</FormLabel>
            <Input type="text" placeholder="imfoobar" />
          </FormControl>
          <FormControl isRequired mb="4">
            <FormLabel>Email</FormLabel>
            <Input type="email" placeholder="foobar@live.me" />
          </FormControl>
          <FormControl isRequired mb="4">
            <FormLabel>Password</FormLabel>
            <Input type="password" placeholder="********" />
          </FormControl>
          <FormControl isRequired mb="8">
            <FormLabel>Repeat Password</FormLabel>
            <Input type="password" placeholder="********" />
          </FormControl>
          <Button colorScheme="blue" px="8" mx="auto" display="block" mb="4">
            Sign Up
          </Button>
          <Text align="center" color="gray.600">
            Already have an account?
            <Link color="blue.600"> Login here</Link>
          </Text>
        </Box>
      </Center>
    </Box>
  );
}
