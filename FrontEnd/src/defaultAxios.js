import axios from "axios";

const getToken = () => localStorage.getItem("token");

const defaultAxios = axios.create({
  baseURL: "http://localhost:8080", // Ensure this is the correct backend URL
});

defaultAxios.interceptors.request.use((config) => {
  const token = getToken();
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export { defaultAxios };
