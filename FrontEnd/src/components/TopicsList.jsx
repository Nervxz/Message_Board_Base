import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { defaultAxios } from "../defaultAxios";
import "../index.css";

const TopicsList = () => {
  const [topics, setTopics] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchTopics = async () => {
      try {
        const response = await defaultAxios.get("/topics/");
        setTopics(response.data);
        setLoading(false);
      } catch (err) {
        setError(err.message);
        setLoading(false);
      }
    };

    fetchTopics();
  }, []);

  const handleUpvote = async (topicId) => {
    try {
      await defaultAxios.post(`/topics/${topicId}/upvote`);
      const response = await defaultAxios.get("/topics/");
      setTopics(response.data);
    } catch (err) {
      setError(err.message);
    }
  };

  if (loading) {
    return <p>Loading topics...</p>;
  }

  if (error) {
    return <p>Error loading topics: {error}</p>;
  }

  return (
    <div className="topics-list max-w-4xl mx-auto mt-10 p-6 bg-white shadow-md rounded-md border border-gray-300">
      <h2 className="text-2xl font-bold mb-4 mt-10 text-center">Topics</h2>
      {topics.length === 0 ? (
        <p>No topics available.</p>
      ) : (
        <ul>
          {topics.map((topic) => (
            <li
              key={topic.TopicID}
              className="mb-10 p-4 mt-5 bg-gray-100 shadow-md rounded-md border border-black"
            >
              <Link to={`/topic/${topic.TopicID}`} className="text-blue-500">
                <h3 className="text-xl font-bold">{topic.Title}</h3>
              </Link>
              <p>{topic.Body}</p>
              <p className="text-sm text-gray-500">
                Published on: {new Date(topic.DatePublished).toLocaleString()}
              </p>
              <div className="flex items-center mt-2">
                <button
                  onClick={() => handleUpvote(topic.TopicID)}
                  className="mr-2 px-4 py-2 bg-blue-500 text-white rounded"
                >
                  Upvote
                </button>
                <span>{topic.Upvotes} Upvotes</span>
              </div>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default TopicsList;
