import {
  Box,
  Button,
  Center,
  FormErrorMessage,
  FormControl,
  FormLabel,
  Heading,
  Input,
  Link,
  Text,
} from "@chakra-ui/react";
import React from "react";
import { useForm } from "react-hook-form";

export default function Login(): JSX.Element {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();
  const onSubmit = (data: Record<string, unknown>) => console.log(data);
  const onError = (err: Record<string, unknown>) => console.log(err);

  return (
    <Box py="10rem">
      <Center>
        <Box
          onSubmit={handleSubmit(onSubmit, onError)}
          as="form"
          minW="lg"
          p="6"
          rounded="md"
          bgColor="white"
          shadow="md"
        >
          <Heading color="gray.800" textAlign="center" mb="2">
            Login
          </Heading>
          <Text color="gray.600" align="center" mb="8">
            You need to login first before posting stuff
          </Text>
          <FormControl isRequired mb="4">
            <FormLabel>Email</FormLabel>
            <Input
              {...register("email", {
                required: "required",
                pattern: {
                  value: /\S+@\S+\.\S+/,
                  message: "Please enter a valid email!",
                },
              })}
              type="text"
              autoComplete="off"
              placeholder="foobar@live.me"
            />
            <Text>{errors?.email?.message}</Text>
            {/* <FormErrorMessage></FormErrorMessage> */}
          </FormControl>
          <FormControl isRequired mb="8">
            <FormLabel>Password</FormLabel>
            <Input
              {...register("password", { minLength: 8, required: "required" })}
              type="password"
              placeholder="********"
            />
          </FormControl>
          <Button
            colorScheme="blue"
            px="8"
            mx="auto"
            mb="4"
            display="block"
            type="submit"
          >
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
