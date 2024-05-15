import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Navbar from "./components/Navbar";
import SignIn from "./components/SignIn";
import SignUp from "./components/SignUp";

const App = () => {
  return (
    <Router>
      <Navbar />
      <div className="mt-16 p-">
        {" "}
        {/* Add margin top to avoid content overlapping with the navbar */}
        <Routes>
          <Route path="/signin" element={<SignIn />} />
          <Route path="/signup" element={<SignUp />} />

          {/* Add other routes here */}
        </Routes>
      </div>
    </Router>
  );
};

export default App;
