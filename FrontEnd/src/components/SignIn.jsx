import React, { useState } from "react";
import { defaultAxios } from "../defaultAxios";
import { useAuth } from "../context/AuthContext";
import { useNavigate } from "react-router-dom";
import "../index.css";

const SignIn = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");
  const { signIn } = useAuth();
  const navigate = useNavigate();

  // Handle sign in
  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await defaultAxios.post("/auth/signin", {
        username,
        password,
      });
      setMessage("Sign in successful");
      console.log("Sign in successful", response.data);
      signIn(response.data.token);
      navigate("/"); // Redirect to home page
    } catch (error) {
      setMessage("Error signing in");
      console.error("Error signing in", error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <h2 className="text-2xl font-bold mb-4">Sign In</h2>
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
      <button type="submit">Sign In</button>
      {message && <p>{message}</p>}
    </form>
  );
};

export default SignIn;
