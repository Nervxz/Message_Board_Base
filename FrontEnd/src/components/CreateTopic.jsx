import React, { useState } from "react";
import { useAuth } from "../context/AuthContext";
import { useNavigate } from "react-router-dom";
import "../index.css";
import { defaultAxios } from "../defaultAxios";

const CreateTopic = () => {
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const { token } = useAuth();
  const [message, setMessage] = useState("");
  const navigate = useNavigate();

  // Handle create topic
  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await defaultAxios.post(
        "/topics/",
        { Title: title, Body: description },
        { headers: { Authorization: `Bearer ${token}` }, withCredentials: true }
      );
      setMessage("Topic created successfully");
      console.log("Topic created successfully", response.data);
      // Redirect to home page after 1 seconds
      setTimeout(() => {
        navigate("/");
      }, 1000);
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
      <h2 className="text-2xl font-bold mb-4 text-center">Create Topics</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-4">
          <label className="block text-gray-700">Title:</label>
          <input
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            className="w-full px-3 py-2 border rounded-md focus:ring focus:border-blue-300"
          />
        </div>
        <div className="mb-4">
          <label className="block text-gray-700">Description:</label>
          <textarea
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            className="w-full px-3 py-2 border rounded-md focus:outline-none focus:ring focus:border-blue-300"
          ></textarea>
        </div>
        <button
          type="submit"
          className="w-full bg-blue-500 text-black py-2 rounded-md hover:bg-blue-600"
        >
          Create Topic
        </button>
        {message && <p className="mt-4 text-center text-red-500">{message}</p>}
      </form>
    </div>
  );
};

export default CreateTopic;
