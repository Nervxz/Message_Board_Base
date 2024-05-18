import React, { useEffect, useState } from "react";
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
              className="mb-10 p-4 mt-10 bg-gray-100 shadow-md rounded-md border border-black"
            >
              <h3 className="text-xl font-bold">{topic.Title}</h3>

              <p className="text-sm text-gray-500">
                Published on: {new Date(topic.DatePublished).toLocaleString()}
              </p>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default TopicsList;
