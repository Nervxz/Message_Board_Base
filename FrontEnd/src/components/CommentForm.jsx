import React, { useState } from "react";
import { defaultAxios } from "../defaultAxios";
import { useAuth } from "../context/AuthContext";

// eslint-disable-next-line react/prop-types
const CommentForm = ({ topicID }) => {
  const [comment, setComment] = useState("");
  const { token } = useAuth();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      // Include the token in the headers
      const response = await defaultAxios.post(
        "/comments/",
        {
          Comment: comment,
          TopicID: parseInt(topicID, 10), // Convert the topicID to an integer because topicID could be a string, (base-10) for using digits 0-9
        },
        {
          headers: {
            Authorization: `Bearer ${token}`, // Send the token in the header
          },
          withCredentials: true,
        }
      );

      console.log("Comment posted successfully:", response.data);
      setComment("");
    } catch (error) {
      console.error("Error posting comment:", error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <textarea
        value={comment}
        onChange={(e) => setComment(e.target.value)}
        placeholder="Write your comment here"
      />
      <button type="submit">Submit Comment</button>
    </form>
  );
};

export default CommentForm;
