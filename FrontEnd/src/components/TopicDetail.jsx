import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { defaultAxios } from "../defaultAxios";
import CommentForm from "./CommentForm";

const TopicDetail = () => {
  const { id } = useParams();
  const [topic, setTopic] = useState(null);
  const [comments, setComments] = useState([]); // Initialize as an empty array
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
    return <p>{error}</p>;
  }

  return (
    <div className="container mx-auto p-4">
      {topic && (
        <div>
          <h1 className="text-2xl font-bold">{topic.Title}</h1>
          <p>{topic.Body}</p>
        </div>
      )}
      <div>
        <h2 className="text-xl font-bold mt-4">Comments</h2>
        {comments.length > 0 ? (
          comments.map((comment) => (
            <div key={comment.CommentID} className="border p-2 my-2">
              <p>{comment.Comment}</p>
            </div>
          ))
        ) : (
          <p>No comments yet.</p>
        )}
      </div>
      <CommentForm topicID={id} />
    </div>
  );
};

export default TopicDetail;
