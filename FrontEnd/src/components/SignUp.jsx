import React, { useState } from "react";
import { defaultAxios } from "../defaultAxios";
import { useNavigate } from "react-router-dom";

const SignUp = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");
  const navigate = useNavigate();
  // Handle sign up
  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await defaultAxios.post("/auth/signup", {
        username,
        password,
      });
      setMessage("Sign up successful");
      console.log("Sign up successful", response.data);
      navigate("/signin"); // Redirect to SignIn page
    } catch (error) {
      setMessage("Error signing up");
      console.error("Error signing up", error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <h2 className="text-2xl font-bold mb-4">Sign Up</h2>
        <label>Username:</label>
        <input
          className="border border-sky-500"
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
      </div>
      <div>
        <label>Password:</label>
        <input
          className="border border-sky-500"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
      </div>
      <button type="submit">Sign Up</button>
      {message && <p>{message}</p>}
    </form>
  );
};

export default SignUp;
