import React, { useState } from "react";
import { defaultAxios } from "../defaultAxios";
import { useNavigate } from "react-router-dom";
import "../index.css";

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
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <form
        onSubmit={handleSubmit}
        className="w-full max-w-md p-8 bg-white rounded shadow-md"
      >
        <h2 className="text-2xl font-bold mb-6 text-center">Sign Up</h2>
        <div className="mb-4">
          <label className="block text-gray-700">Username:</label>
          <input
            className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-sky-500"
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>
        <div className="mb-6">
          <label className="block text-gray-700">Password:</label>
          <input
            className="w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-sky-500"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <button
          type="submit"
          className="w-full py-2 bg-sky-500 text-white font-bold rounded hover:bg-sky-600 focus:outline-none focus:ring-2 focus:ring-sky-500"
        >
          Sign Up
        </button>
        {message && <p className="mt-4 text-center text-red-500">{message}</p>}
      </form>
    </div>
  );
};

export default SignUp;
