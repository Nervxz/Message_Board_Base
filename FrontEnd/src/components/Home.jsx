import React from "react";
import Navbar from "./Navbar";
import TopicsList from "./TopicsList";

const Home = () => {
  return (
    <>
      <Navbar />
      <div className="home-page flex flex-col items-center justify-start min-h-screen bg-gray-100 pt-20">
        {" "}
        {/* Add padding-top to move content below Navbar */}
        <TopicsList />
      </div>
    </>
  );
};

export default Home;
