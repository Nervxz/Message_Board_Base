import axios from "axios";

const instance = axios.create({
  baseURL: "http://localhost:8080",
  withCredentials: true, // Include credentials in requests
});

export default instance;
