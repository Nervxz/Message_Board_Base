import React, { useState } from "react";
import { defaultAxios } from "../defaultAxios";

const CreateTopic = () => {
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");
  const [message, setMessage] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const sessionToken = localStorage.getItem("session_token"); // Retrieve the token from local storage
      if (!sessionToken) {
        throw new Error("No session token found. Please log in.");
      }

      const response = await defaultAxios.post(
        "/topics/",
        { Title: title, Body: body }, // No need to send UserID, it will be extracted from the session token on the backend
        { headers: { Authorization: `Bearer ${sessionToken}` } } // Include the token in the Authorization header
      );
      setMessage("Topic created successfully");
      console.log("Topic created successfully", response.data);
    } catch (error) {
      setMessage("Error creating topic");
      console.error(
        "Error creating topic:",
        error.response ? error.response.data : error.message
      );
    }
  };

  return (
    <div className="max-w-md mx-auto mt-10 p-6 bg-white shadow-md rounded-md">
      <h2 className="text-2xl font-bold mb-4 text-center">Create Topic</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-4">
          <label className="block text-gray-700">Title:</label>
          <input
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring focus:border-blue-300"
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Body:</label>
          <textarea
            value={body}
            onChange={(e) => setBody(e.target.value)}
            className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring focus:border-blue-300"
          ></textarea>
        </div>
        <button
          type="submit"
          className="w-full bg-blue-500 text-white py-2 rounded-md hover:bg-blue-600"
        >
          Create Topic
        </button>
        {message && <p className="mt-4 text-center text-red-500">{message}</p>}
      </form>
    </div>
  );
};

export default CreateTopic;
