import React, { useState } from "react";
import axios from "../axiosConfig"; // Import the configured Axios instance

const SignIn = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post("/auth/signin", { username, password });
      setMessage("Sign in successful");
      console.log("Sign in successful", response.data);
      // Save the token in local storage or state management
      localStorage.setItem("token", response.data.token);
    } catch (error) {
      setMessage("Error signing in");
      console.error("Error signing in", error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Username:</label>
        <input
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
      </div>
      <div>
        <label>Password:</label>
        <input
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
      </div>
      <button type="submit">Sign In</button>
      {message && <p>{message}</p>}
    </form>
  );
};

export default SignIn;
