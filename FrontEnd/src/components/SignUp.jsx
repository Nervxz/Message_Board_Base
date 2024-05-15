import React, { useState } from "react";
import axios from "axios";

const SignUp = () => {
  const [Username, setUsername] = useState("");
  const [Password, setPassword] = useState("");
  const [message, setMessage] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post("http://localhost:8080/auth/signup", {
        Username,
        Password,
      });
      setMessage("Sign up successful");
      console.log("Sign up successful", response.data);
    } catch (error) {
      setMessage("Error signing up");
      console.error("Error signing up", error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Username:</label>
        <input
          type="text"
          value={Username}
          onChange={(e) => setUsername(e.target.value)}
        />
      </div>
      <div>
        <label>Password:</label>
        <input
          type="password"
          value={Password}
          onChange={(e) => setPassword(e.target.value)}
        />
      </div>
      <button type="submit">Sign Up</button>
      {message && <p>{message}</p>}
    </form>
  );
};

export default SignUp;
