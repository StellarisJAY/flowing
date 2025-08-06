import axios from "axios";

const http = axios.create({
    baseURL: import.meta.env.VITE_API_URL,
    timeout: 10000,
});

http.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem("flowing_access_token");
        if (token) {
            config.headers['X-Access-Token'] = token;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);
http.interceptors.response.use(
  (response) => {
    const {code, message} = response.data;
    if (code === 200) {
      return response.data;
    }
    if (code === 401) {
      localStorage.removeItem("flowing_access_token");
      window.location.href = "/sys/login";
      return;
    }
    return Promise.reject(message || "Error");
  },
  (error) => {
    console.log(error);
  }
);

export default http;
