import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { defaultAxios } from "../defaultAxios";
import CommentForm from "./CommentForm";

const TopicDetail = () => {
  const { id } = useParams();
  const [topic, setTopic] = useState(null);
  const [comments, setComments] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchTopic = async () => {
      try {
        const response = await defaultAxios.get(`/topics/${id}`);
        setTopic(response.data);
      } catch (error) {
        console.error("Error fetching topic:", error);
        setError("Error fetching topic");
      }
    };

    const fetchComments = async () => {
      try {
        const response = await defaultAxios.get(`/comments/${id}`);
        setComments(response.data);
      } catch (error) {
        console.error("Error fetching comments:", error);
        setError("Error fetching comments");
      }
    };

    fetchTopic();
    fetchComments();
  }, [id]);

  if (error) {
    return <p className="text-red-500">{error}</p>;
  }

  return (
    <div className="container mx-auto p-6 bg-gray-100 min-h-screen">
      {topic ? (
        <div className="bg-white shadow-lg rounded-lg p-8 mb-6">
          <h1 className="text-4xl font-bold text-gray-900 mb-4">
            {topic.Title}
          </h1>
          <p className="text-gray-700 text-lg">{topic.Body}</p>
        </div>
      ) : (
        <p className="text-center text-gray-500">Loading topic...</p>
      )}
      <div className="bg-white shadow-lg rounded-lg p-8">
        <h2 className="text-3xl font-bold text-gray-900 mb-4">Comments</h2>
        {comments.length > 0 ? (
          comments.map((comment) => (
            <div
              key={comment.CommentID}
              className="border-b border-gray-200 py-4"
            >
              <p className="text-gray-800">{comment.Comment}</p>
            </div>
          ))
        ) : (
          <p className="text-gray-500">No comments yet.</p>
        )}
      </div>
      <div className="mt-6">
        <CommentForm topicID={id} />
      </div>
    </div>
  );
};

export default TopicDetail;
