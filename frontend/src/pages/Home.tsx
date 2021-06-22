import { Box, Center, Container, Heading, Text } from "@chakra-ui/react";
import React from "react";

export default function Home(): JSX.Element {
  return (
    <Box pt="10rem">
      <Center>
        <Container maxW="xl" textAlign="center">
          <Heading as="h1" color="gray.800" mb="4">
            Welcome to Post-It!
          </Heading>
          <Text fontSize="xl" color="gray.600" lineHeight="tall">
            It&apos;s basically vaporware at this point but nobody cares anyway
            lol. There are still lots of stuff that needs to be implemented.
          </Text>
        </Container>
      </Center>
    </Box>
  );
}
