import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { defaultAxios } from "../defaultAxios";
import CommentForm from "./CommentForm";

const TopicDetail = () => {
  const { id } = useParams();
  const [topic, setTopic] = useState(null);
  const [comments, setComments] = useState([]);

  useEffect(() => {
    const fetchTopic = async () => {
      try {
        const response = await defaultAxios.get(`/topics/${id}`);
        setTopic(response.data);
      } catch (error) {
        console.error("Error fetching topic:", error);
      }
    };

    const fetchComments = async () => {
      try {
        const response = await defaultAxios.get(`/comments/${id}`);
        setComments(response.data);
      } catch (error) {
        console.error("Error fetching comments:", error);
      }
    };

    fetchTopic();
    fetchComments();
  }, [id]);

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
        {comments.map((comment) => (
          <div key={comment.CommentID} className="border p-2 my-2">
            <p>{comment.Comment}</p>
          </div>
        ))}
      </div>
      <CommentForm topicID={id} />
    </div>
  );
};

export default TopicDetail;
