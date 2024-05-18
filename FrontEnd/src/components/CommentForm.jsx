import React, { useState } from "react";
import { defaultAxios } from "../defaultAxios";

const CommentForm = ({ topicID }) => {
  const [comment, setComment] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await defaultAxios.post(`/comments`, {
        Comment: comment,
        TopicID: topicID,
      });
      setComment("");
      // Optionally refresh comments
    } catch (error) {
      console.error("Error submitting comment:", error);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="mt-4">
      <textarea
        value={comment}
        onChange={(e) => setComment(e.target.value)}
        className="w-full border p-2"
        placeholder="Add a comment"
        required
      ></textarea>
      <button type="submit" className="bg-blue-500 text-white px-4 py-2 mt-2">
        Submit
      </button>
    </form>
  );
};

export default CommentForm;
