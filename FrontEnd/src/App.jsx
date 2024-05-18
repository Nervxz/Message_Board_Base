import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import SignIn from "./components/SignIn";
import SignUp from "./components/SignUp";
import Home from "./components/Home";
import CreateTopic from "./components/CreateTopic";
import ProtectedRoute from "./components/ProtectedRoute";
import TopicDetail from "./components/TopicDetail"; // Import the TopicDetail component
import { AuthProvider } from "./context/AuthContext";
import "./index.css";

const App = () => {
  return (
    <AuthProvider>
      <Router>
        <Routes>
          <Route path="/signin" element={<SignIn />} />
          <Route path="/signup" element={<SignUp />} />
          <Route path="/" element={<Home />} />
          <Route path="/topics/:id" element={<TopicDetail />} />{" "}
          {/* Add this line */}
          <Route
            path="/create-topic"
            element={
              <ProtectedRoute>
                <CreateTopic />
              </ProtectedRoute>
            }
          />
        </Routes>
      </Router>
    </AuthProvider>
  );
};

export default App;
