import React from "react";
import Navbar from "./Navbar";
import TopicsList from "./TopicsList";
import "../index.css";

const Home = () => {
  return (
    <>
      <Navbar />
      <div className="home-page mt-100 flex flex-col items-center justify-start min-h-screen">
        <TopicsList />
      </div>
    </>
  );
};

export default Home;
