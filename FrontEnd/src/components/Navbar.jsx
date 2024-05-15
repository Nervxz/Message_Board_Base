import React from "react";
import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <nav className="bg-gray-800 p-4 fixed top-0 w-full z-10">
      <div className="container mx-auto flex justify-between items-center">
        <div className="text-white text-lg">
          <Link to="/">Home</Link>
        </div>
        <div className="text-white flex space-x-4 ml-auto">
          {" "}
          {/* Added ml-auto to move links to the right */}
          <Link to="/signin" className="px-4">
            Sign In
          </Link>
          <Link to="/signup" className="px-4">
            Sign Up
          </Link>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
