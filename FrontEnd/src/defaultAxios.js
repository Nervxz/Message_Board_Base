import axios from "axios";

const defaultAxios = axios.create({
  baseURL: "http://localhost:8080", // Ensure this is the correct backend URL
});

export { defaultAxios };
