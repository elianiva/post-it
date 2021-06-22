import React from "react";
import { BrowserRouter as Router, Route } from "react-router-dom";
import Login from "./pages/Login";

function App(): JSX.Element {
  return (
    <Router>
      <Route path="/login" component={Login} />
    </Router>
  );
}

export default App;
