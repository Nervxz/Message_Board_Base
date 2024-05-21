import React, { useState } from "react";
import { defaultAxios } from "../defaultAxios";
import { useAuth } from "../context/AuthContext";

// eslint-disable-next-line react/prop-types
const CommentForm = ({ topicID }) => {
  const [comment, setComment] = useState("");
  const { token } = useAuth();
  const [error, setError] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await defaultAxios.post(
        "/comments/",
        {
          Comment: comment,
          TopicID: parseInt(topicID, 10),
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
          withCredentials: true,
        }
      );

      console.log("Comment posted successfully:", response.data);
      setComment("");
      setError(null);
    } catch (error) {
      if (error.response && error.response.status === 401) {
        alert("Sign in to comment");
      } else {
        console.error("Error posting comment:", error);
        setError("Error posting comment");
      }
    }
  };

  return (
    <div className="bg-white shadow-md rounded-lg p-6 mt-6">
      <h3 className="text-2xl font-bold mb-4">Leave a Comment</h3>
      {error && <p className="text-red-500 mb-4">{error}</p>}
      <form onSubmit={handleSubmit}>
        <textarea
          value={comment}
          onChange={(e) => setComment(e.target.value)}
          placeholder="Write your comment here"
          className="w-full p-4 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 mb-4"
          rows="4"
        />
        <button
          type="submit"
          className="bg-blue-500 text-white font-bold py-2 px-4 rounded-lg hover:bg-blue-600 transition duration-300"
        >
          Submit Comment
        </button>
      </form>
    </div>
  );
};

export default CommentForm;
