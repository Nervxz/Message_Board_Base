// CommentForm.jsx
import React, { useState } from "react";
import { defaultAxios } from "../defaultAxios";

// eslint-disable-next-line react/prop-types
const CommentForm = ({ topicId }) => {
  const [comment, setComment] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      // Retrieve the token from localStorage
      const token = localStorage.getItem("token");

      // Include the token in the headers
      const response = await defaultAxios.post(
        "/comments/",
        {
          comment,
          topicId,
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
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
        placeholder="Write your comment..."
        required
      />
      <button type="submit">Submit</button>
    </form>
  );
};

export default CommentForm;
