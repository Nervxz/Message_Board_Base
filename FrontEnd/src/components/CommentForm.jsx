import React, { useState } from "react";
import { defaultAxios } from "../defaultAxios";
import { useAuth } from "../context/AuthContext";

// eslint-disable-next-line react/prop-types
const CommentForm = ({ topicId }) => {
  const [comment, setComment] = useState("");
  const { token } = useAuth();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
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
        placeholder="Write your comment..."
        required
      />
      <button type="submit">Submit</button>
    </form>
  );
};

export default CommentForm;
