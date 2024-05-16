import axios from "axios";
const getToken = () => localStorage.getItem("token"); // Assuming the token is stored in localStorage

const defaultAxios = axios.create({
  baseURL: "http://localhost:8080", // Ensure this is the correct backend URL
  headers: {
    Authorization: `Bearer ${getToken()}`, // Include the token in the Authorization header
  },
});

export { defaultAxios };
